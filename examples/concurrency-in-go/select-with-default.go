package main

import (
	"fmt"
	"time"
)

func main() {
	// const totalTime = time.Second * 1
	const totalTime = time.Microsecond * 1

	done := make(chan struct{})

	go func() {
		time.Sleep(totalTime)
		close(done)
	}()

	incCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		incCount++
		// time.Sleep(1 * time.Second)
	}
	fmt.Printf("incCount: %v\n", incCount)
	fmt.Printf("Achieved %.3e cycles of work before signalled to stop.\n",
		float64(incCount))
}
