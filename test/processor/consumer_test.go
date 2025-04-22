package main

import (
	"context"
	"testing"
	"time"

	"github.com/aprofessionale/crypto-alert-system/internal/common"
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

type MockKafkaReader struct {
	messages []kafka.Message
}

func (m *MockKafkaReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	if len(m.messages) == 0 {
		<-ctx.Done()
		return kafka.Message{}, ctx.Err()
	}
	msg := m.messages[0]
	m.messages = m.messages[1:]
	return msg, nil
}

func TestConsumePrices(t *testing.T) {
	mockMsg := kafka.Message{
		Key:   []byte("BTC"),
		Value: []byte(`{"Symbol":"BTC","Price":61234.56,"Time":1713400000}`),
	}

	reader := &MockKafkaReader{
		messages: []kafka.Message{mockMsg},
	}

	received := make(chan common.PriceData, 1)

	go func() {
		err := processor.ConsumePrices(reader, received)
		if err != nil {
			t.Errorf("Consumer prices failed: %v ", err)
		}
	}()

	select {
	case p := <-received:
		if p.Symbol != "BTC" || p.Price != 61234.56 {
			t.Errorf("Unexpected price data %+v", p)
		}
	case <-time.After(2 * time.Second):
		t.Error("did not receive message.")
	}
}
