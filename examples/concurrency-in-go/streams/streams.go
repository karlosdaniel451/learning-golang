package streams

type Stream[T any] <-chan T

func (stream Stream[T]) Collect() []T {
	collectedValues := []T{}

	for value := range stream {
		collectedValues = append(collectedValues, value)
	}

	return collectedValues
}

// Mapped every value using `mapper` and return a new Stream with such mapped
// values.
func (stream Stream[T]) Map(
	done <-chan struct{},
	mapper func(T) T,
) Stream[T] {

	mappedStream := make(chan T)

	go func() {
		defer close(mappedStream)

		for value := range stream {
			select {
			case <-done:
				return
			default:
			}

			mappedStream <- mapper(value)
		}
	}()

	return mappedStream
}

// Filter values according to a predicate and return a new stream with such
// filtered values.
func (stream Stream[T]) Filter(
	done <-chan struct{},
	predicate func(T) bool,
) Stream[T] {

	filteredStream := make(chan T)

	go func() {
		defer close(filteredStream)

		for value := range stream {
			select {
			case <-done:
				return
			default:
			}

			if predicate(value) {
				filteredStream <- value
			}
		}
	}()

	return filteredStream
}

// Returns a new stream with only the first `n` values from `stream`.
func (stream Stream[T]) Take(
	done <-chan struct{},
	n int,
) Stream[T] {

	takedStream := make(chan T)

	go func() {
		defer close(takedStream)

		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			default:
			}

			takedValue := <-stream
			takedStream <- takedValue
		}
	}()

	return takedStream
}
