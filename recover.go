package main

import (
	"fmt"
	"log"
)

func printDivision(a, b int) {
	defer func() {
		if p := recover(); p != nil {
			log.Println(p)
		}
	}()

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func betterPrintDivision(a, b int) error {
	if b == 0 {
		return fmt.Errorf("impossible division by zero: %d / %d", a, b)
	}

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	return nil
}

func main() {
	for i := 5; i >= 0; i-- {
		printDivision(5, i)
	}

	for i := 50; i >= 0; i -= 10 {
		printDivision(50, i)
	}

	for i := 5; i >= 0; i-- {
		if err := betterPrintDivision(5, i); err != nil {
			log.Println(err)
		}
	}

	for i := 50; i >= 0; i -= 10 {
		if err := betterPrintDivision(50, i); err != nil {
			log.Println(err)
		}
	}
}
