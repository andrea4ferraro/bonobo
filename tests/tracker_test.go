package tests

import (
	"testing"

	"liquidity-position-tracker/internal/tracker"
)

func TestSample(t *testing.T) {

	p := tracker.Sample()

	if p.Pool == "" {

		t.Fail()
	}
}
