package main

import (
	"context"
	"testing"
)

func TestPublishPrice(t *testing.T) {
	price := PriceData{
		Symbol: "BTC",
		Price:  60200.12,
		Time:   1713100000,
	}

	err := PublishPrice(context.Background(), price)

	if err != nil {
		t.Errorf("failed to publish price: %v", err)
	}
}
