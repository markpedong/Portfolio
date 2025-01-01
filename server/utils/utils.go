package utils

import (
	"strings"
	"unicode"
)

func SanitizeFileName(fileName string) string {
	// Replace spaces with underscores
	fileName = strings.ReplaceAll(fileName, " ", "_")

	// Remove special characters and keep only alphanumeric characters, underscores, and periods
	var sanitized []rune
	for _, r := range fileName {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' {
			sanitized = append(sanitized, r)
		}
	}

	return strings.ToLower(string(sanitized))
}
