package main

import (
	"fmt"
	"time"
)

func main() {
	n := 45

	cancelSignal := make(chan struct{})
	go spinner(100*time.Millisecond, cancelSignal)

	fibonacci := fibonacci(n)
	cancelSignal <- struct{}{}

	fmt.Printf("fib(%d) = %d\n", n, fibonacci)
}

func spinner(delay time.Duration, cancelSignal <-chan struct{}) {
	for {
		select {
		case <-cancelSignal:
			eraseOutputLine()
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(delay)
			}
		}
	}
}

func eraseOutputLine() {
	// Thanks ChatGPT for making me find this mysterious string sequence to
	// erase the current printed output:
	fmt.Print("\033[2K\r")
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
