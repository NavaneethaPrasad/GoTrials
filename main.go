package main

import (
	"fmt"
)

func main() {

	fmt.Println(fizzbizz(15))

	fmt.Println(palindrome("hai"))
	fmt.Println(palindrome("Mom"))

	fmt.Println(abbreviate("Indian Institute of Management"))
	fmt.Println(abbreviate("Automated Teller Machine"))

	fmt.Println(pangram("The quick brown fox jumps over the lazy dog"))
	fmt.Println(pangram("I am Navaneetha"))

	fmt.Println(frequency("Bobby"))
	fmt.Println(frequency("hai"))
}
