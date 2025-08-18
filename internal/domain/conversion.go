// currency-converter/internal/domain/conversion.go
package domain

import "time"

// Conversion representa uma única transação de conversão.
type Conversion struct {
	ID              int       `json:"id"`
	FromCurrency    string    `json:"from_currency"`
	ToCurrency      string    `json:"to_currency"`
	InitialAmount   float64   `json:"initial_amount"`
	ConvertedAmount float64   `json:"converted_amount"`
	Rate            float64   `json:"rate"`
	CreatedAt       time.Time `json:"created_at"`
}

// ConversionRequest é o que a API espera receber no corpo da requisição.
type ConversionRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
