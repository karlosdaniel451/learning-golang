package main

import (
	"fmt"
	"sync"
)

type Element struct {
	value int
	mutex sync.Locker
}

func main() {
	element := Element{value: 0, mutex: &sync.RWMutex{}}

	var wg sync.WaitGroup

	const numberOfIncrements int = 10000
	const numberOfDecrements int = 10000

	for i := 0; i < numberOfIncrements; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment(&element)
		}()
	}

	for i := 0; i < numberOfDecrements; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement(&element)
		}()
	}

	wg.Wait()

	fmt.Printf("\nelement value: %d\n", element.value)
}

func increment(element *Element) {
	element.mutex.Lock()
	defer element.mutex.Unlock()
	element.value++
	// fmt.Printf("Incremented from %d to %d\n", element.value-1, element.value)
	fmt.Print("++ ")
}

func decrement(element *Element) {
	element.mutex.Lock()
	defer element.mutex.Unlock()
	element.value--
	// fmt.Printf("Decremented from %d to %d\n", element.value+1, element.value)
	fmt.Print("-- ")
}
