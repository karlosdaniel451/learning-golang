package main

import (
	"errors"
	"fmt"
	"log"
)

func MapInts(ints []int, mapper func(int) int) {
	for i := range ints {
		ints[i] = mapper(ints[i])
	}
}

func ReduceInts(ints []int, reducer func(int, int) int) (int, error) {

	if len(ints) == 0 {
		return 0, errors.New("no int value passed")

	}

	if len(ints) == 1 {
		return ints[0], nil
	}

	reducedValue := ints[0]

	for i := 1; i < len(ints); i++ {
		reducedValue = reducer(ints[i], reducedValue)
	}

	return reducedValue, nil
}

func FilterInts(ints []int, condition func(int) bool) []int {
	newSlice := make([]int, 0, len(ints))

	for i := range ints {
		if condition(ints[i]) == true {
			newSlice = append(newSlice, ints[i])
		}
	}

	return newSlice
}

func main() {
	integers := []int{1, 1, 3, 5}
	fmt.Println(integers)

	MapInts(integers, func(integer int) int { return integer * 2 })
	fmt.Println(integers)

	sum, err := ReduceInts(integers, func(a, b int) int { return a + b })
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)

	fmt.Println(FilterInts(integers, func(x int) bool { return x%2 != 0 }))
}
