package main

import "fmt"

func buildIncrementer(start, step int) func() int {
	x := start

	return func() int {
		x += step
		return x
	}
}

func main() {
	oneStepIncrementer := buildIncrementer(0, 1)
	fmt.Println(oneStepIncrementer())
	fmt.Println(oneStepIncrementer())
	fmt.Println(oneStepIncrementer())
	fmt.Println()

	twoStepIncrementer := buildIncrementer(2, 2)
	fmt.Println(twoStepIncrementer())
	fmt.Println(twoStepIncrementer())
	fmt.Println(twoStepIncrementer())
	fmt.Println()

	threeStepIncrementer := buildIncrementer(1, 3)
	fmt.Println(threeStepIncrementer())
	fmt.Println(threeStepIncrementer())
	fmt.Println(threeStepIncrementer())
	fmt.Println()

	fmt.Println(buildIncrementer(0, 10)())
}
