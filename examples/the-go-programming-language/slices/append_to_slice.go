package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var words []string

	for {
		fmt.Print("Write a word or \\stop to stop: ")
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		word = strings.TrimRight(word, "\n")

		if word == `\stop` {
			break
		}

		words = append(words, word)
	}

	fmt.Printf("\nNumber of read words: %d\n", len(words))
	fmt.Printf("Read words: [%s]\n", strings.Join(words, ", "))
}
