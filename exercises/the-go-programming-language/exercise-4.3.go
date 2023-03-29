/*
Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
*/

package main

import "fmt"

func main() {
	var array [100]int
	for i := range array {
		array[i] = i
	}

	fmt.Println(array)

	reverseArray(&array)

	fmt.Println(array)
}

func reverseArray(arrayPtr *[100]int) {
	j := len(*arrayPtr) - 1
	for i := 0; i < len(*arrayPtr)/2; i++ {
		arrayPtr[i], arrayPtr[j] = arrayPtr[j], arrayPtr[i]
		j--
	}
}
