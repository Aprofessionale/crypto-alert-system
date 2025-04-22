package collector_test

import (
	"testing"

	"github.com/aprofessionale/crypto-alert-system/cmd/collector"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMockPrice(t *testing.T) {
	price := collector.GenerateMockPrice()

	assert.NotEmpty(t, price.Symbol)
	assert.True(t, price.Price > 0)
}
