package tests

import (
	"testing"

	"liquidity-position-tracker/internal/calculator"
)

func TestProfit(t *testing.T) {

	if calculator.Profit(150, 100) != 50 {

		t.Fail()
	}
}
