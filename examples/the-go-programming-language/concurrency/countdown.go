package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)

	fmt.Println("Commencing countdown. Press return to abort.")

	go func() {
		os.Stdin.Read(make([]byte, 1)) // Block until read a single byte
		abort <- struct{}{}
	}()

	abortionTime := countdownLaunch(tick, abort)

	if !abortionTime.IsZero() {
		fmt.Printf("Lanching aborted at %s\n", abortionTime.UTC())
		return
	}
	launch()
}

func countdownLaunch(tick <-chan time.Time, abort <-chan struct{}) (abortionTime time.Time) {
	for countdown := 10; countdown >= 0; countdown-- {
		select {
		case <-tick:
			if countdown == 0 {
				break
			}
			fmt.Println(countdown)
		case <-abort:
			return time.Now().UTC()
		}
	}

	return time.Time{}
}

func launch() {
	fmt.Println("Launching rocket!!!")
}
