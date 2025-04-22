package collector

import (
	"math/rand"
	"time"
)

type PriceData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Time   int64   `json:"time"`
}

func GenerateMockPrice() PriceData {
	symbols := []string{"BTC", "ETH"}
	symbol := symbols[rand.Intn(len(symbols))]

	basePrice := map[string]float64{
		"BTC": 60000,
		"ETH": 3000,
	}

	price := basePrice[symbol] + rand.Float64()*1000 - 500 // +/- 500
	return PriceData{
		Symbol: symbol,
		Price:  price,
		Time:   time.Now().Unix(),
	}
}
