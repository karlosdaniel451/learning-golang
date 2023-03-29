package main

import "fmt"

func main() {
	var numbers [10]int
	for i := range numbers {
		numbers[i] = i + 1
	}

	multiplyArray(&numbers, 2)

	for i, value := range numbers {
		fmt.Println(i, value)
	}
}

func multiplyArray(arrayPtr *[10]int, factor int) {
	for i := range arrayPtr {
		arrayPtr[i] *= factor
	}
}
