package utils

import "fmt"

func USD(v float64) string {

	return fmt.Sprintf("$%.2f", v)
}
