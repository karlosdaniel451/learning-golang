package main

import "fmt"

func max(integers ...int) int {
	if len(integers) == 0 {
		return 0
	}

	currentMax := integers[0]
	for _, integer := range integers {
		if integer > currentMax {
			currentMax = integer
		}
	}
	return currentMax
}

func min(integers ...int) int {
	if len(integers) == 0 {
		return 0
	}

	currentMin := integers[0]
	for _, integer := range integers {
		if integer < currentMin {
			currentMin = integer
		}
	}
	return currentMin
}



func main() {
	fmt.Println(max())	
	fmt.Println(min())	

	fmt.Println(max(5))	
	fmt.Println(min(5))	

	fmt.Println(max(5, 0, -5))	
	fmt.Println(min(5, 0, -5))	

	fmt.Println(max([]int{5, 0, -5}...))
	fmt.Println(min([]int{5, 0, -5}...))
}
