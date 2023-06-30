package streams

import "sync"

// Multiplex multiple Streams into a single one without preserving
// the order of the incoming values.
func FanIn[T any](done <-chan struct{}, streams ...Stream[T]) Stream[T] {
	multiplexedStream := make(chan T)

	var wg sync.WaitGroup

	multiplex := func(stream Stream[T]) {
		defer wg.Done()

		for value := range stream {
			select {
			case <-done:
				return	
			default:
			}

			multiplexedStream <- value
		}
	}

	// Multiplex all Streams.
	for _, stream := range streams {
		wg.Add(1)
		go multiplex(stream)
	}

	// Wait for all Streams to be drained..
	go func() {
		wg.Wait()	
		close(multiplexedStream)
	}()

	return multiplexedStream
}
