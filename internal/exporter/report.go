package exporter

import (
	"liquidity-position-tracker/internal/storage"
)

func Export(v any) {

	storage.Save("data/report.json", v)
}
