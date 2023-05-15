package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	// Equivalent implementation:
	// for {
	// 	i, ok := <- in
	// 	if !ok {
	// 		break
	// 	}
	// 	out <- i * i
	// }
	// close(out)
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
