/*
Exercise 1.2: Modify the echo program to print the index and value of each
of its arguments, one per line.
*/
package main

import (
	"fmt"
	"os"
)

func exercise_1_7_main() {
	for i, arg := range os.Args {
		fmt.Println(i, "-", arg)
	}
}
