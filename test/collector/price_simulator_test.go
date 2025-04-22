package collector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMockPrice(t *testing.T) {
	price := GenerateMockPrice()

	assert.NotEmpty(t, price.Symbol)
	assert.True(t, price.Price > 0)
}
