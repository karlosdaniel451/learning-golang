package main

import (
	"log"
	"time"
)

func main() {

	cancellation1 := setInterval(func() {
		log.Println("hello 2s")
	}, time.Second*2)

	cancellation2 := setInterval(func() {
		log.Println("hello 4s")
	}, time.Second*4)

	cancellation3 := setInterval(func() {
		log.Println("hello 8s")
	}, time.Second*8)

	time.Sleep(time.Second * 30)

	cancellation1 <- struct{}{}
	cancellation2 <- struct{}{}
	cancellation3 <- struct{}{}

	// close(cancellation1)
	// close(cancellation2)
	// close(cancellation3)

	log.Println("timeouts closed")
	time.Sleep(time.Second * 10)
}

func setInterval(fn func(), delay time.Duration) (cancelChannel chan struct{}) {
	goroutineStarted := make(chan struct{})
	cancelChannel = make(chan struct{})
	go func() {
		goroutineStarted <- struct{}{}
		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		// for {
		// 	select {
		// 	// Check for stop condition
		// 	case <-cancelChannel:
		// 		return
		// 	case <-ticker.C:
		// 		fn()
		// 	}
		// }

		for {
			select {
			case <-cancelChannel:
				return
			default:
				break
			}

			select {
			case <-ticker.C:
				fn()
			default:
				break
			}
		}

	}()
	<-goroutineStarted
	return cancelChannel
}
