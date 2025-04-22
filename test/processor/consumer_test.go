package main

import (
	"context"
	"testing"
	"time"

	processor "github.com/aprofessionale/crypto-alert-system/internal/processor"
	"github.com/segmentio/kafka-go"
)

func TestConsumePriceMessage(t *testing.T) {
	// Send a test message to Kafka first
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "crypto-prices",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	msg := kafka.Message{
		Key:   []byte("BTC"),
		Value: []byte(`{"Symbol":"BTC","Price":12345.67,"Time":1713100000}`),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		t.Fatalf("Failed to write test message: %v", err)
	}

	// Start consumer in a goroutine
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		err := processor.ConsumePriceMessages(ctx)
		if err != nil {
			t.Errorf("consumer error: %v", err)
		}
	}()

	time.Sleep(2 * time.Second) // Allow time to consume

	// You could add logs/assertions or mock the handler later
}
