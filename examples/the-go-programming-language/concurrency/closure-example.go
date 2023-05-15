package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Microsecond)
		fmt.Println("Printing in a separate goroutine...")
	}()

	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println("Printing while another goroutine is executing...")
	}
	fmt.Printf("Printed after %s\n", time.Since(start))
}
