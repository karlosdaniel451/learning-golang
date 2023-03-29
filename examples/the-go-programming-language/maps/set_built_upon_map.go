package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read a line from stdin and write it to stdout if the line wasn't read before
func main() {
	reader := bufio.NewReader(os.Stdin)
	lines := make(map[string]bool)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		line = strings.TrimRight(line, "\n")

		if !lines[line] {
			lines[line] = true
			fmt.Println(line)
		}
		fmt.Println()
	}
}
