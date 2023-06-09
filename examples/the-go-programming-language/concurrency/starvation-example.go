package main

import "fmt"

func main() {
	// starvationExample1()
	// starvationExample2()
}

func starvationExample1() {
	const channelBufferCapacity int = 5
	ch := make(chan int, channelBufferCapacity)

	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	// By here, the buffered channel is full, and any send operation
	// will be blocked until a read operation be made. Since there is
	// no other goroutine, there will be a deadlock because the main
	// goroutine will wail forever.

	ch <- 1 // Wait forever here.
}

func starvationExample2() {
	const channelBufferCapacity int = 5
	ch := make(chan int, channelBufferCapacity)

	// By here, the buffered channel is empty, and any read operation
	// will be blocked until a send operation be made. Since there is
	// no other goroutine, there will be a deadlock because the main
	// goroutine will wail forever.

	fmt.Println(<-ch)
}
