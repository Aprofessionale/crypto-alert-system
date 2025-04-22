package collector

import (
	"math/rand"
	"time"

	common "github.com/aprofessionale/crypto-alert-system/internal/common"
)

func GenerateMockPrice() common.PriceData {
	symbols := []string{"BTC", "ETH"}
	symbol := symbols[rand.Intn(len(symbols))]

	basePrice := map[string]float64{
		"BTC": 60000,
		"ETH": 3000,
	}

	price := basePrice[symbol] + rand.Float64()*1000 - 500 // +/- 500
	return common.PriceData{
		Symbol: symbol,
		Price:  price,
		Time:   time.Now().Unix(),
	}
}
