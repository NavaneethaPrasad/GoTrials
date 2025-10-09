package main

import (
	"fmt"
	"strings"
)

func main() {
	var c1 celsiues = 100.0
	fmt.Printf("Celsius :%f,Farenheit:%f \n", c1, C2F(c1))
	s := "golang"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c:%d %T \n", s[i], s[i], s[i])
	}
	if strings.Contains(s, "la") {
		fmt.Printf(s, "contains x")
	} else {
		fmt.Printf(s, "Does not contain x")
	}

	fmt.Println(s)

	fmt.Println(fizzbizz(15))

	fmt.Println(palindrome("hai"))
	fmt.Println(palindrome("mom"))

	fmt.Println(abbreviate("Indian Institute of Management"))
	fmt.Println(abbreviate("Automated Teller Machine"))

	fmt.Println(pangram("The quick brown fox jumps over the lazy "))

	fmt.Println(frequency("Bobby"))
}
