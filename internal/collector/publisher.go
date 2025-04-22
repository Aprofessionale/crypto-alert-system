package collector

import (
	"context"
	"encoding/json"
	"log"

	common "github.com/aprofessionale/crypto-alert-system/internal/common"
	"github.com/segmentio/kafka-go"
)

func PublishPrice(ctx context.Context, price common.PriceData) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "crypto-prices",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	value, err := json.Marshal(price)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(price.Symbol),
		Value: value,
	})
	if err != nil {
		log.Printf("Failed to write to Kafka: %v", err)
		return err
	}

	log.Printf("Publisher price: %s -> %f", price.Symbol, price.Price)
	return nil
}
