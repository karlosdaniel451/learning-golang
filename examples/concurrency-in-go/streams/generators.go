package streams

func StreamGenerator[T any](done <-chan struct{}, values ...T) Stream[T] {
	generatedStream := make(chan T)

	go func() {
		defer close(generatedStream)

		for _, value := range values {
			select {
			case <-done:
				return
			default:
			}

			generatedStream <- value
		}
	}()

	return generatedStream

}

// Repeatedly calls a function and return a Stream with the generated
// values on an as-needed basis.
func RepeatFunctionGenerator[T any](done <-chan struct{}, fn func() T) Stream[T] {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			default:
			}

			stream <- fn()
		}
	}()

	return stream
}

func RepeatStreamGenerator[T any](done <-chan struct{}, values ...T) Stream[T] {
	generatedStream := make(chan T)

	go func() {
		defer close(generatedStream)

		for {
			for _, value := range values {
				select {
				case <-done:
					return
				default:
				}

				generatedStream <- value
			}
		}
	}()

	return generatedStream
}
