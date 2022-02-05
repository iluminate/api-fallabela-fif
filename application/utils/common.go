package utils

import "strconv"

func ParseFloat(amountStr string) float64 {
	amountInt, _ := strconv.ParseFloat(amountStr, 64)
	return amountInt
}
