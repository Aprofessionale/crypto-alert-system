package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	common "github.com/aprofessionale/crypto-alert-system/internal/common"
	"github.com/segmentio/kafka-go"
)

type KafkaReader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

func ConsumePriceMessages(ctx context.Context) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "crypto-prices",
		GroupID: "processor-group",
	})
	defer reader.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			var price common.PriceData
			if err := json.Unmarshal(m.Value, &price); err != nil {
				log.Printf("Invalid JSON: %s", m.Value)
				continue
			}

			fmt.Printf("Received price: %s - $%.2f\n", price.Symbol, price.Price)
		}
	}
}

func ConsumePrices(reader KafkaReader, out chan<- common.PriceData) error {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var price common.PriceData
		if err := json.Unmarshal(msg.Value, &price); err != nil {
			log.Printf("Failed to unmarshal: %v ", err)
			continue
		}

		log.Printf("Consumed: %+v", price)
		out <- price
	}
}
