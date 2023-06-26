package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hi 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hi 2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hi 3")
	}()

	wg.Wait()
}
