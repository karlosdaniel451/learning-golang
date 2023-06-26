package main

import (
	"fmt"

	"slices"
)

func main() {
	const listMaxLength int = 10

	list := make([]int, 0, listMaxLength)

	for i := 0; i < listMaxLength; i++ {
		list = append(list, i)
	}

	fmt.Printf("list: %v\n", list)

	list = slices.Delete[[]int](list, 0, 1)

	fmt.Printf("list: %v\n", list)

	list = slices.Insert[[]int](list, 0, []int{-1}...)

	fmt.Printf("list: %v\n", list)
}
