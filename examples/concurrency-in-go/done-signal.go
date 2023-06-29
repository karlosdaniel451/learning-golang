package main

import (
	"fmt"
	"time"
)

func main() {
	const producerDelay time.Duration = time.Second * 1
	const consumerDelay time.Duration = time.Second * 1

	stringStream := make(chan string)

	done := make(chan struct{})

	producerTerminated := produceStrings(done, stringStream, producerDelay)
	consumerTerminated := consumeStrings(done, stringStream, consumerDelay)

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("Canceling producer and consumer...")
		close(done)
	}()

	<-producerTerminated
	<-consumerTerminated

	fmt.Println("Producer and consumer terminated, exiting program...")

}

func produceStrings(
	done <-chan struct{},
	stringStream chan<- string,
	delay time.Duration,
) (terminated chan struct{}) {

	terminated = make(chan struct{})

	go func() {
		defer fmt.Println("Producer is finished.")
		defer close(terminated)
		for {
			time.Sleep(delay)
			select {
			case <-done:
				return
			default:
			}

			stringStream <- "asdf"
		}
	}()

	return terminated
}

func consumeStrings(
	done <-chan struct{},
	stringStream <-chan string,
	delay time.Duration,
) (terminated chan struct{}) {

	terminated = make(chan struct{})

	go func() {
		defer fmt.Println("Consumer is finished.")
		defer close(terminated)
		for {
			time.Sleep(delay)
			select {
			case <-done:
				return
			default:
			}

			fmt.Println(<-stringStream)
		}
	}()

	return terminated
}
