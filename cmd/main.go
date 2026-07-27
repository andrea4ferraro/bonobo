package main

import (
	"fmt"

	"liquidity-position-tracker/internal/exporter"
	"liquidity-position-tracker/internal/tracker"
)

func main() {

	position := tracker.Sample()

	fmt.Println(position)

	exporter.Export(position)
}
