package main

import (
	"fmt"
)

func main() {
	fmt.Println(exercise2([]string{"hi", "ha", "hi", "bye", "hello", "bye"}))
}

// Write a function that takes a slice of strings and returns a map containing
// the frequency of each word in the slice.
func exercise2(words []string) map[string]int {
	wordsFrequency := make(map[string]int, len(words))

	for _, word := range words {
		wordsFrequency[word]++
	}

	return wordsFrequency
}
