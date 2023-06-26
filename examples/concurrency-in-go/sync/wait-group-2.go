package main

import (
	"fmt"
	"sync"
)

func main() {
	const numberOfGreetings int = 10
	var wg sync.WaitGroup

	wg.Add(numberOfGreetings)

	for i := 0; i < numberOfGreetings; i++ {
		go printHello(&wg, i+1)
	}

	wg.Wait()
}

func printHello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello, %d\n", id)
}
