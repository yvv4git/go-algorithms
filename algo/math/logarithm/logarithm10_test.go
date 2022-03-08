package logarithm

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Десятичный логарифм.
func TestLogarithm10(t *testing.T) {
	t.Log("Test - logarithm 10")

	// Десятичный логарифм.
	// В какую степень надо возвести 10, чтобы получить 10, 100, 1000
	// log10 10 = 1
	// log10 100 = 2
	// log10 1000 = 3
	testCases := []struct {
		InValue  float64
		Expected float64
	}{
		{1, 0},
		{10, 1},
		{100, 2},
		{1000, 3},
		{10000, 4},
		{100000, 5},
	}

	for _, tc := range testCases {
		actualValue := math.Log10(tc.InValue)
		assert.Equal(t, tc.Expected, actualValue)
	}
}
