package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	common "github.com/aprofessionale/crypto-alert-system/internal/common"
	"github.com/segmentio/kafka-go"
)

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
