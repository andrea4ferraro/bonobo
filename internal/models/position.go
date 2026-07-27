package models

type Position struct {

	Pool string `json:"pool"`

	TokenA string `json:"token_a"`

	TokenB string `json:"token_b"`

	AmountA float64 `json:"amount_a"`

	AmountB float64 `json:"amount_b"`

	ValueUSD float64 `json:"value_usd"`

	ProfitUSD float64 `json:"profit_usd"`
}
