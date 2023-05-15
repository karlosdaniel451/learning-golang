package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	executionTime := make(chan time.Duration)
	stopConditionOfFeedback := make(chan struct{})
	fmt.Println("Doing some heavy operation...")

	go func() {
		start := time.Now()
		simulateHeavyOperation()
		executionTime <- time.Since(start)
	}()

	// Execute the following function call concurrently with the heavy operation
	// in order to provide the feedback to the user.
	go printFeedbackOfHeavyOperationExecution(stopConditionOfFeedback)

	// The main goroutine will wait for the goroutine executing the heavy
	// operation to complete.
	executionTimeOfOperation := <-executionTime

	// Signal the feedback writing to be stopped, since the heavy operation
	// was completed.
	stopConditionOfFeedback <- struct{}{}

	eraseOutputLine()
	fmt.Printf("Done after %.3f seconds.\n", executionTimeOfOperation.Seconds())
}

func printFeedbackOfHeavyOperationExecution(stop <-chan struct{}) {
	i := 0
	for {
		select {
		case <-stop:
			return // Stop writing feedback if stop condition is sent.
		default:
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			if i%3 == 2 {
				fmt.Print("\r")
				eraseOutputLine()
			}
			i++
		}
	}
}

/*
Simulate a heavy computation operation. It could be both an I/O bound
operation and a CPU-bound one. This operation is:

- Independent/autonomous: it does not depend on another operation, such as to
receive input data.

- No side effects: it does not change or maintain mutable shared data.

These characteristics make such function an excellent candidate for concurrency
and parellelism without much difficulty.
*/
func simulateHeavyOperation() {
	delayInMs := rand.Intn(10_000)
	time.Sleep(time.Duration(delayInMs) * time.Millisecond)
}

func eraseOutputLine() {
	// Thanks ChatGPT for making me find this mysterious string sequence to
	// erase the current printed output:
	fmt.Print("\033[2K\r")
}
