package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...]string{"John Doe", "", "", "Jane Doe", "Mark Moe", "", "Alice", "Bob"}

	fmt.Println(strings.Join(getNonEmptyStrings(names[:]), ", "))
}

func getNonEmptyStrings(strings []string) []string {
	var nonEmptyStrings []string

	for _, value := range strings {
		if value != "" {
			nonEmptyStrings = append(nonEmptyStrings, value)
		}
	}

	return nonEmptyStrings
}
