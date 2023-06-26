package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("error: %s", fmt.Errorf("insuficient arguments: pass input filename"))
	}

	inputFilename := os.Args[1]

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var currentSumOfAmountsOfCalories int
	var newInput string

	var greatestSum, secondGreatestSum, thirdGreatestSum int

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		newInput = scanner.Text()

		if newInput != "" {
			// In case the calories amount is for the same elf

			newCaloriesAmount, err := strconv.Atoi(newInput)
			if err != nil {
				panic(err)
			}

			currentSumOfAmountsOfCalories += newCaloriesAmount
		} else {
			// In case there is a blank indicating a new elf

			// If current sum is at least the third greatest
			if currentSumOfAmountsOfCalories > thirdGreatestSum {
				greatestSum, secondGreatestSum, thirdGreatestSum = descendingOrderOfTriple(
					greatestSum,
					secondGreatestSum,
					currentSumOfAmountsOfCalories,
				)
			}

			currentSumOfAmountsOfCalories = 0
		}
	}
	// If last current sum is at least the third greatest
	if currentSumOfAmountsOfCalories > thirdGreatestSum {
		greatestSum, secondGreatestSum, thirdGreatestSum = descendingOrderOfTriple(
			greatestSum,
			secondGreatestSum,
			currentSumOfAmountsOfCalories,
		)
	}

	fmt.Println(greatestSum + secondGreatestSum + thirdGreatestSum)
	fmt.Printf("greatestSum: %d\n", greatestSum)
	fmt.Printf("secondGreatestSum: %d\n", secondGreatestSum)
	fmt.Printf("thirdGreatestSum: %d\n", thirdGreatestSum)
}

func descendingOrderOfTriple(a, b, c int) (int, int, int) {
	if a < b {
		a, b = b, a
	}

	if a < c {
		a, c = c, a
	}

	if b < c {
		b, c = c, b
	}

	return a, b, c
}
