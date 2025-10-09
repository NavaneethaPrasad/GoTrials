package main

import (
	"strings"
	"unicode"
)

func abbreviate(s string) string {
	words := strings.Split(s, " ")

	var abbrevation strings.Builder
	for _, word := range words {
		if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
			abbrevation.WriteByte(word[0])
		}
	}
	return abbrevation.String()
}
