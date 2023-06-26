package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	value int
}

func produce(itemStream chan<- *Item, delay time.Duration) {
	for {
		time.Sleep(delay)
		newItem := Item{
			value: rand.Intn(11),
		}
		itemStream <- &newItem
		fmt.Printf(
			"item value produced: %d - buffer size: %d\n",
			newItem.value, len(itemStream),
		)
	}
}

func consume(itemStream <-chan *Item, delay time.Duration) {
	for {
		time.Sleep(delay)
		readItem := <-itemStream
		fmt.Printf(
			"item value consumed: %d - buffer size: %d\n",
			readItem.value, len(itemStream),
		)
	}
}

func main() {
	const (
		numberOfProducers int           = 3
		numberOfConsumers int           = 3
		bufferSize        int           = 10
		producersDelay    time.Duration = time.Second * 1
		consumersDelay    time.Duration = time.Second * 2
	)
	itemStream := make(chan *Item, bufferSize)

	for i := 0; i < numberOfProducers; i++ {
		go produce(itemStream, producersDelay)
	}

	for i := 0; i < numberOfConsumers; i++ {
		go consume(itemStream, consumersDelay)
	}

	time.Sleep(time.Hour)
}
