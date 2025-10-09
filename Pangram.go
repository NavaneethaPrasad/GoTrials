package main

import (
	"strings"
)

func pangram(s string) bool {
	s = strings.ToLower(s)
	for letter := 'a'; letter <= 'z'; letter++ {
		if !strings.Contains(s, string(letter)) {
			return false
		}
	}
	return true
}
