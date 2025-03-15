package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/hamba/avro/v2"
)

func NewProducer(kafkaBrokers string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBrokers})
	if err != nil {
		panic(err)
	}
	return p
}

func deliveryHandler(p *kafka.Producer) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}

func produce(p *kafka.Producer, topic string, generatorFunc func() []interface{}, avroSchema avro.Schema) {
	defer p.Close()

	// Delivery report handler for produced messages
	go deliveryHandler(p)

	// Generate some alerts
	alerts := generatorFunc()

	// Produce messages to topic (asynchronously)
	for i, alert := range alerts {
		marshalled, err := avro.Marshal(avroSchema, alert)
		if err != nil {
			panic(fmt.Errorf("Failed to marshal alert: %w", err))
		}
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          marshalled,
		}, nil)

		// sleep between 0 and 1000 millisecond
		time.Sleep(time.Duration(gofakeit.Number(0, 1000)) * time.Millisecond)
		// after 20 messages, sleep for 2 to 5 seconds
		if i%20 == 0 {
			time.Sleep(time.Duration(gofakeit.Number(1, 5)) * time.Second)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func ParseZTFSchema() (avro.Schema, error) {
	paths := []string{"schemas/candidate.avsc", "schemas/prv_candidate.avsc", "schemas/fp_hist.avsc", "schemas/cutout.avsc", "schemas/alert.avsc"}
	return avro.ParseFiles(paths...)
}

func ParseStampClassifierSchema() (avro.Schema, error) {
	paths := []string{"schemas/early_classification.avsc"}
	return avro.ParseFiles(paths...)
}

func generateZtfAlerts() []any {
	alerts := make([]any, 1000)
	for i := 0; i < len(alerts); i++ {
		var alert ZtfAlert
		err := gofakeit.Struct(&alert)
		if err != nil {
			panic(fmt.Errorf("Failed to generate fake alert: %w", err))
		}
		alert.Candidate.Ra = gofakeit.Latitude()
		alert.Candidate.Dec = gofakeit.Longitude()
		alert.Candidate.Magpsf = gofakeit.Float32Range(10, 25)

		alerts[i] = alert
	}
	return alerts
}

func generateStampClassifierMessages() []any {
	messages := make([]any, 1000)
	for i := 0; i < len(messages); i++ {
		var message StampProbabilities
		err := gofakeit.Struct(&message)
		if err != nil {
			panic(fmt.Errorf("Failed to generate fake alert: %w", err))
		}
		messages[i] = message
	}
	return messages
}
