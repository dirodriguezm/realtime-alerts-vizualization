package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	var cmd string
	var kafkaBrokers string
	var mock bool
	flag.StringVar(&cmd, "cmd", "server", "Command to run: server or producer")
	flag.StringVar(&kafkaBrokers, "kafka-brokers", "localhost:9092", "Kafka brokers to connect to")
	flag.BoolVar(&mock, "mock", false, "Use mock object repository")
	flag.Parse()

	var username *string
	var password *string
	env := os.Getenv("KAFKA_USERNAME")
	if env != "" {
		username = &env
	}
	env = os.Getenv("KAFKA_PASSWORD")
	if env != "" {
		password = &env
	}

	topic := "stamp_classifier_20250314"
	var objectRepository ObjectRepository
	objectRepository = ALeRCEObjectRepository{}
	if mock {
		slog.Info("Using mock object repository")
		objectRepository = MockObjectRepository{}
	}
	switch cmd {
	case "server":
		start(kafkaBrokers, topic, objectRepository, username, password)
	case "producer":
		schema, err := ParseStampClassifierSchema()
		if err != nil {
			panic(err)
		}
		producer := NewProducer(kafkaBrokers)
		produce(producer, topic, generateStampClassifierMessages, schema)
	default:
		panic("Invalid command")
	}
}
