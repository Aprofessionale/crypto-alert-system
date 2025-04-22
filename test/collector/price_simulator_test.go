package collector_test

import (
	"testing"

	collector "github.com/aprofessionale/crypto-alert-system/internal/collector"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMockPrice(t *testing.T) {
	price := collector.GenerateMockPrice()

	assert.NotEmpty(t, price.Symbol)
	assert.True(t, price.Price > 0)
}
