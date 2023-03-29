package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, err := convertSliceFromStringToInt(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input to int: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Original values: %v\n", numbers)

	reverseSlice(numbers)

	fmt.Printf("Reversed values: %v\n", numbers)
}

func convertSliceFromStringToInt(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))

	for i := range intSlice {
		parsedInt, err := strconv.ParseInt(stringSlice[i], 10, 0)
		if err != nil {
			return nil, err
		}
		intSlice[i] = int(parsedInt)
	}

	return intSlice, nil
}

func reverseSlice(slice []int) {
	j := len(slice) - 1
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}
}
