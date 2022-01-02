package utils

import "strconv"

func ParseFloat(amountStr string, amountDefault float64) float64 {
	var amountInt float64
	if amountStr != "" {
		amountInt, _ = strconv.ParseFloat(amountStr, 64)
	} else {
		amountInt = amountDefault
	}
	return amountInt
}
