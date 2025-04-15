package main

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func produceTestMessage() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "crypto-prices",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("test"),
			Value: []byte("hello from collector"),
		},
	)
	if err != nil {
		log.Fatal("failed to write message: ", err)
	}
	log.Println("test message sent successfully")
}
