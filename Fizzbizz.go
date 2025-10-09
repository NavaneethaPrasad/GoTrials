package main

import (
	"strconv"
)

func fizzbizz(n int) []string {
	var fb []string
	for m := 1; m < n+1; m++ {
		if m%15 == 0 {
			fb = append(fb, "fizzbizz")
		} else if m%5 == 0 {
			fb = append(fb, "bizz")
		} else if m%3 == 0 {
			fb = append(fb, "fizz")
		} else {
			fb = append(fb, strconv.Itoa(m))
		}
	}
	return fb
}
