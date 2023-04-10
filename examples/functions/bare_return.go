package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	integers, err := ParseStringsToInt(os.Args[1:])
	if err != nil {
		// fmt.Fprintf(os.Stderr, "%v\n", err)
		// os.Exit(1)
		log.Fatal(err)
	}

	even, odd := CountEvenAndOdd(integers)

	// fmt.Printf("count of even numbers: %d\ncount of odd numbers: %d\n", even, odd)
	fmt.Printf("quantidade de números pares: %d\nquantidade de números ímpares: %d\n", even, odd)
}

func CountEvenAndOdd(numbers []int) (even, odd int) {
	if len(numbers) == 0 {
		return // It is equivalent to `return even, odd` and it is called "bare return"
	}

	for _, number := range numbers {
		if number%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return // It is equivalent to `return even, odd` and it is called "bare return"
}

func ParseStringsToInt(stringSlice []string) ([]int, error) {
	// The following is more efficient than `integers := []int{}` since it
	// does not require additional memory allocations.
	integers := make([]int, 0, len(stringSlice))

	for _, stringElement := range stringSlice {
		parsedInteger, err := strconv.Atoi(stringElement)
		if err != nil {
			// return nil, fmt.Errorf("error by parsing \"%s\" as integer", stringElement)
			return nil, fmt.Errorf("erro ao converter \"%s\" como inteiro", stringElement)
		}

		integers = append(integers, parsedInteger)
	}

	return integers, nil
}
