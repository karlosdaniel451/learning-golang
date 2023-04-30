package main

import "fmt"

type Counter struct {
	// The Counter value cannot be set to arbitraty values, only incremented or reset,
	// that is why it is encapsulated. Thus, the `value` field is not exported to other
	// packages.
	value int
}

func (counter *Counter) Value() int {
	return counter.value
}

func (counter *Counter) Increment() {
	counter.value++
}

func (counter *Counter) Reset() {
	counter.value = 0
}

func main() {
	counter := Counter{}
	fmt.Println(counter.value)

	counter.Increment()
	fmt.Println(counter.value)

	counter.Reset()
	fmt.Println(counter.value)

	counter.Increment()
	counter.Increment()
	fmt.Println(counter.value)
}
