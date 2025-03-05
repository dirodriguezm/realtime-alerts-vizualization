package main

import (
	"flag"
)

func main() {
	var cmd string
	var kafkaBrokers string
	flag.StringVar(&cmd, "cmd", "server", "Command to run: server or producer")
	flag.StringVar(&kafkaBrokers, "kafka-brokers", "localhost:9092", "Kafka brokers to connect to")
	flag.Parse()

	switch cmd {
	case "server":
		start(kafkaBrokers)
	case "producer":
		produce(kafkaBrokers)
	default:
		panic("Invalid command")
	}
}
