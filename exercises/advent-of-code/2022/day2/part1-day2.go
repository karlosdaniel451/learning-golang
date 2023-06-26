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

	var greatestSumOfCalories, currentSumOfAmountsOfCalories int
	var newInput string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		newInput = scanner.Text()

		if newInput != "" { // In case the calories amount is for the same elf
			newCaloriesAmount, err := strconv.Atoi(newInput)
			if err != nil {
				panic(err)
			}

			currentSumOfAmountsOfCalories += newCaloriesAmount
		} else { // In case there is the break line indicating a new elf
			if currentSumOfAmountsOfCalories > greatestSumOfCalories {
				greatestSumOfCalories = currentSumOfAmountsOfCalories
			}

			currentSumOfAmountsOfCalories = 0
		}
	}

	fmt.Println(greatestSumOfCalories)
}
