package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do " +
		"eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut " +
		"enim ad minim veniam, quis nostrud exercitation ullamco laboris " +
		"nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in " +
		"reprehenderit in voluptate velit esse cillum dolore eu fugiat " +
		"nulla pariatur. Excepteur sint occaecat cupidatat non proident, " +
		"sunt in culpa qui officia deserunt mollit anim id est laborum."

	fmt.Println(exercise3(text))
}

// Write a function that takes a map containing the frequency of each word in
// a text and a word and returns the frequency of the word in the text.
func exercise3(text string) map[string]int {
	// Clean text removing unnecessary characters.
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, ",", " ")

	// Split text producing the words slice.
	words := strings.Split(text, " ")

	wordsFrequency := make(map[string]int, len(words))

	for _, word := range words {
		wordsFrequency[word]++
	}

	return wordsFrequency
}
