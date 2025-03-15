package main

import (
	"log/slog"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/hamba/avro/v2"
)

func NewConsumer(servers, groupId, topic string, username, password *string) *kafka.Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  servers,
		"group.id":           groupId,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": "false",
	}

	if username != nil && password != nil {
		slog.Info("Using SASL authentication")
		config.SetKey("sasl.mechanisms", "SCRAM-SHA-512")
		config.SetKey("security.protocol", "SASL_SSL")
		config.SetKey("sasl.username", *username)
		config.SetKey("sasl.password", *password)
	}

	c, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		panic(err)
	}

	slog.Info("Created connection to kafka", "consumer", c)

	return c
}

func consume(ch chan StampProbabilities, c *kafka.Consumer) {
	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			alert := StampProbabilities{}
			avroSchema, err := ParseStampClassifierSchema()
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
			_, commitErr := c.Commit()
			if commitErr != nil {
				slog.Error("Failed to commit message", "error", err)
			}
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
