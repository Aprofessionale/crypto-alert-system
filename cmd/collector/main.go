package main

import (
	"context"
	"time"

	collector "github.com/aprofessionale/crypto-alert-system/internal/collector"
)

func main() {
	for {
		price := collector.GenerateMockPrice()
		collector.PublishPrice(context.Background(), price)
		time.Sleep(2 * time.Second)
	}
}
