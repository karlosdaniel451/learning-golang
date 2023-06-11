/*
4.22 Write a multithreaded program that calculates various statistical values
for a list of numbers. This program will be passed a series of numbers
on the command line and will then create three separate worker threads.
One thread will determine the average of the numbers, the second will
determine the maximum value, and the third will determine the mini-
mum value. For example, suppose your program is passed the integers:
	90 81 78 95 79 72 85
The program will report:
	The average value is 82
	The minimum value is 72
	The maximum value is 95
*/

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Invalid input: insuficient arguments. Pass the start, stop and step values.")
	}
	start, stop, step, err := parseSequenceInput(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	sequence := generateSequence(start, stop, step)

	// if numbers != nil {
	// 	fmt.Println(numbers)
	// }

	var average float64
	var min, max int
	reversedSequence := make([]int, len(sequence))

	var wg sync.WaitGroup

	startTime := time.Now()
	wg.Add(1)
	// func() {
	go func() {
		defer wg.Done()
		fmt.Println("Computing average...")
		average = getAverage(sequence)
		fmt.Println("Average computed.")
	}()

	wg.Add(1)
	// func() {
	go func() {
		defer wg.Done()

		fmt.Println("Computing minimum...")
		min = getMin(sequence)
		fmt.Println("Minimum computed.")
	}()

	wg.Add(1)
	// func() {
	go func() {
		defer wg.Done()

		fmt.Println("Computing maximum...")
		max = getMax(sequence)
		fmt.Println("Maximum computed.")
	}()

	wg.Add(1)
	// func() {
	go func() {
		defer wg.Done()

		fmt.Println("Sorting sequence...")
		sort.Slice(sequence, func(i, j int) bool { return sequence[i] < sequence[j] })
		fmt.Println("Sequence sorted...")
	}()

	wg.Add(1)
	// func() {
	go func() {
		defer wg.Done()

		fmt.Println("Reversing sequence...")
		copy(reversedSequence, sequence)
		sort.Slice(reversedSequence, func(i, j int) bool { return reversedSequence[i] > reversedSequence[j] })
		fmt.Println("Sequence reversed...")
	}()

	wg.Wait()


	fmt.Printf(
		"Average: %.2f\nMinimum: %d\nMaximum: %d\nSorted: %v...\nReversed: %v...\n",
		average, min, max, sequence[0:5], reversedSequence[0:5],
	)

	fmt.Printf("Computing done after %.3f seconds.\n", time.Since(startTime).Seconds())
}

func generateSequence(start, stop, step int) []int {
	sequence := make([]int, 0, (stop-start)/step)

	for i := start; i < stop; i += step {
		sequence = append(sequence, i)
	}

	return sequence
}

func parseSequenceInput(start, stop, step string) (int, int, int, error) {
	parsedStart, err := strconv.Atoi(start)
	if err != nil {
		return 0, 0, 0, err
	}
	parsedStop, err := strconv.Atoi(stop)
	if err != nil {
		return 0, 0, 0, err
	}
	parsedStep, err := strconv.Atoi(step)
	if err != nil {
		return 0, 0, 0, err
	}

	return parsedStart, parsedStop, parsedStep, nil
}

func getAverage(numbers []int) float64 {
	return float64(getSum(numbers) / len(numbers))
}

func getMax(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	max := numbers[0]

	for _, number := range numbers[1:] {
		if number > max {
			max = number
		}
	}

	return max
}

func getMin(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	min := numbers[0]

	for _, number := range numbers[1:] {
		if number < min {
			min = number
		}
	}

	return min
}
func getSum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
