package tracker

import "liquidity-position-tracker/internal/models"

func Sample() models.Position {

	return models.Position{

		Pool: "ETH/USDC",

		TokenA: "ETH",

		TokenB: "USDC",

		AmountA: 0.45,

		AmountB: 780,

		ValueUSD: 1820,

		ProfitUSD: 120,
	}
}
