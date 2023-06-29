package main

type IntStream <-chan int

func IntAdder(additive int) func(intStreamValue int) int {
	return func(intStreamValue int) int {
		return additive + intStreamValue
	}
}

func IntStreamGenerator(done <-chan struct{}, ints ...int) IntStream {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for _, integer := range ints {
			select {
			case <-done:
				return
			default:
			}

			intStream <- integer
		}
	}()

	return intStream
}

func MapIntStream(
	done <-chan struct{},
	mapper func(int) int,
	intStream <-chan int,
) <-chan int {

	outputStream := make(chan int)

	go func() {
		defer close(outputStream)

		for integer := range intStream {
			select {
			case <-done:
				return
			default:
			}

			outputStream <- mapper(integer)
		}
	}()

	return outputStream
}

func (inputIntStream IntStream) Map(
	done <-chan struct{},
	mapper func(int) int,
) IntStream {

	outputIntStream := make(chan int)

	go func() {
		defer close(outputIntStream)

		for inputValue := range inputIntStream {
			select {
			case <-done:
				return
			default:
			}

			outputValue := mapper(inputValue)
			outputIntStream <- outputValue
		}
	}()

	return outputIntStream
}
