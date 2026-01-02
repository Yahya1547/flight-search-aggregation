package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatIDR(amount float64) string {
	formatter := message.NewPrinter(language.Indonesian)
	return "Rp " + formatter.Sprintf("%.0f", amount)
}