package collector_test

import (
	"context"
	"testing"

	"github.com/aprofessionale/crypto-alert-system/cmd/collector"
)

func TestPublishPrice(t *testing.T) {
	price := collector.PriceData{
		Symbol: "BTC",
		Price:  60200.12,
		Time:   1713100000,
	}

	err := collector.PublishPrice(context.Background(), price)

	if err != nil {
		t.Errorf("failed to publish price: %v", err)
	}
}
