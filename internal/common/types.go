package common

type PriceData struct {
	Symbol string  `json:"Symbol"`
	Price  float64 `json:"Price"`
	Time   int64   `json:"Time"`
}
