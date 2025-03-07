package main

import (
	"log/slog"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"github.com/hamba/avro/v2"
)

func consume(ch chan ZtfAlert, kafkaBrokers string) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          uuid.New().String(),
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"ztf_alerts"}, nil)

	if err != nil {
		panic(err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			alert := ZtfAlert{}
			avroSchema, err := parseSchema()
			if err != nil {
				slog.Error("Failed to parse schema", "error", err)
				break
			}

			err = avro.Unmarshal(avroSchema, msg.Value, &alert)
			if err != nil {
				slog.Error("Failed to unmarshal alert", "error", err)
				break
			}

			ch <- alert
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			slog.Error("Consumer error", "error", err, "message", msg)
			break
		}

	}

	c.Close()
}
