package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/hamba/avro/v2"
)

func produce(kafkaBrokers string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBrokers})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
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
	}()

	avroSchema, err := parseSchema()
	if err != nil {
		panic(fmt.Errorf("Failed to parse schema: %w", err))
	}

	// Generate some alerts
	alerts := make([]ZtfAlert, 1000)
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

	// Produce messages to topic (asynchronously)
	topic := "ztf_alerts"
	for _, alert := range alerts {
		marshalled, err := avro.Marshal(avroSchema, alert)
		if err != nil {
			panic(fmt.Errorf("Failed to marshal alert: %w", err))
		}
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          marshalled,
		}, nil)
		time.Sleep(2 * time.Second)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func parseSchema() (avro.Schema, error) {
	paths := []string{"schemas/candidate.avsc", "schemas/prv_candidate.avsc", "schemas/fp_hist.avsc", "schemas/cutout.avsc", "schemas/alert.avsc"}
	return avro.ParseFiles(paths...)
}
