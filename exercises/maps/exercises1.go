package main

import (
	"fmt"
	"sort"
)

func main() {
	exercise1CharacterFrequency := exercise1("aabcdddeéeeãeefÀÀàççxwyzß")
	exercise1Characters := []string{}
	for character := range exercise1CharacterFrequency {
		exercise1Characters = append(exercise1Characters, character)
	}

	fmt.Printf("exercise1Characters: %v\n", exercise1Characters)

	sort.Strings(exercise1Characters)

	for _, character := range exercise1Characters {
		fmt.Printf("%s: %d\n", character, exercise1CharacterFrequency[character])
	}
}

// Write a function that takes a string and returns a map containing the
// frequency of each character in the string.
func exercise1(s string) map[string]int {
	characterFrequency := make(map[string]int)

	for _, character := range s {
		characterFrequency[string(character)]++
	}

	return characterFrequency
}

