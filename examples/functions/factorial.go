package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		// fmt.Fprintf(os.Stderr, "error: value for n was not passed\n")
		// os.Exit(1)
		log.Fatal("error: value for n was not passed")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("error by parsing \"%s\" as integer", os.Args[1])
	}

	fmt.Printf("%d! = %d\n", n, factorial(n))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
