package main

import "fmt"

func main() {
	integers := []int{0, 1, 2, 3, 4, 5}

	done1 := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})
	done4 := make(chan struct{})

	intStream1 := IntStreamGenerator(make(<-chan struct{}), integers...)
	intStream2 := IntStreamGenerator(make(<-chan struct{}), integers...)
	intStream3 := IntStreamGenerator(make(<-chan struct{}), integers...)

	square := func(x int) int {
		return x * x
	}

	var outputStream IntStream

	fmt.Println("Original values:")
	for _, integer := range integers {
		fmt.Println(integer)
	}
	fmt.Println("")

	outputStream = intStream1.Map(done1, square)
	fmt.Println("Values to the power of 2 (square):")
	for value := range outputStream {
		fmt.Println(value)
	}
	fmt.Println("")

	outputStream = intStream2.Map(done2, square).Map(done2, square)
	fmt.Println("Values to the power of 4:")
	for value := range outputStream {
		fmt.Println(value)
	}
	fmt.Println("")

	outputStream = intStream3.Map(done3, square).
		Map(done3, square).
		Map(done3, square)

	fmt.Println("Values to the power of 8:")
	for value := range outputStream {
		fmt.Println(value)
	}
	fmt.Println("")

	outputStream = IntStreamGenerator(make(<-chan struct{}), integers...).
		Map(done4, IntAdder(10)).
		Map(done4, IntAdder(-5))

	fmt.Println("Values added by 10 and then subtracted by 5:")
	for value := range outputStream {
		fmt.Println(value)
	}

}
