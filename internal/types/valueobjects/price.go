package valueobjects

import (
	"math"
	"strings"
)

type Price struct {
	Amount   float64
	Currency string
}

func NewPrice(amount float64, currency string) Price {
	return Price{
		Amount:   math.Round(amount*100) / 100,
		Currency: strings.ToUpper(strings.TrimSpace(currency)),
	}
}
