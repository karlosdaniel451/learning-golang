package main

import (
	"fmt"
	"os"
	"strings"
)

func example_echo_main() {
	output := ""
	for _, arg := range os.Args[1:] {
		output += arg + " "
	}
	print(output, "\n")

	fmt.Println(strings.Join(os.Args[1:], " - "))

	fmt.Println(os.Args[1:])
}
