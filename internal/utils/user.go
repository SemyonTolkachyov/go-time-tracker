package utils

import (
	"go-time-tracker/internal/apperror"
	"strings"
)

// ParsePassportNumber get passport series and passport number from combined passport number
func ParsePassportNumber(passportNumber string) (passportSeries string, passportNum string, err error) {
	res := strings.Split(passportNumber, " ")
	if len(res) < 2 {
		return "", "", apperror.NewParseError("invalid passport number")
	}
	return res[0], res[1], nil
}
