package main

func frequency(s string) map[string]int {
	map1 := make(map[string]int)
	for letter := 0; letter < len(s); letter++ {
		ch := string(s[letter])
		map1[ch]++
	}
	return map1
}
