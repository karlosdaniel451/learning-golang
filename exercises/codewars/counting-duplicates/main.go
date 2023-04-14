package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	count := 0

	s1 = strings.ToLower(s1)
	charactersCount := make(map[rune]int)

	for _, character := range s1 {
		charactersCount[character]++
		if charactersCount[character] == 2 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count("aabbcde"))
	fmt.Println(duplicate_count("aabBcde"))
	fmt.Println(duplicate_count("indivisibility"))
	fmt.Println(duplicate_count("Indivisibilities"))
	fmt.Println(duplicate_count("aA11"))
	fmt.Println(duplicate_count("ABBA"))
}
