package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"streams"
	"strings"
	"time"
)

func main() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	var inputStringStream1 streams.Stream[string]
	var inputStringStream2 streams.Stream[string]
	var inputStringStream3 streams.Stream[string]
	var inputStringStream4 streams.Stream[string]

	var outputStringStream streams.Stream[string]

	stringsSlice := []string{
		"a",
		"ab",
		"aab",
		"aaabc",
		"Aaabccc",
	}

	inputStringStream1 = streams.StreamGenerator[string](done1, stringsSlice...)
	inputStringStream2 = streams.StreamGenerator[string](done1, stringsSlice...)
	inputStringStream3 = streams.StreamGenerator[string](done1, stringsSlice...)
	inputStringStream4 = streams.StreamGenerator[string](done2, stringsSlice...)

	fmt.Println("\nOriginal values:")
	for value := range inputStringStream1 {
		fmt.Println(value)
	}

	outputStringStream = inputStringStream2.Map(done1, strings.ToUpper)

	fmt.Println("\nValues in uppper case:")
	for value := range outputStringStream {
		fmt.Println(value)
	}

	outputStringStream = inputStringStream3.Filter(done1, func(s string) bool {
		return strings.Contains(s, "aa")
	})

	fmt.Println("\nValues that contains the \"aa\" substring:")
	for value := range outputStringStream {
		fmt.Println(value)
	}

	const delay1 time.Duration = time.Microsecond * 150

	outputStringStream = streams.RepeatStreamGenerator[string](done1, stringsSlice...)

	time.AfterFunc(delay1, func() {
		close(done1)
	})

	fmt.Printf("\nValues repeated indefinitely for %s:\n", delay1)
	for value := range outputStringStream {
		fmt.Println(value)
	}

	outputStringStream = inputStringStream4.Take(done2, 3)
	fmt.Println("\nFirst 3 values:")
	for value := range outputStringStream {
		fmt.Println(value)
	}

	const delay2 time.Duration = time.Microsecond * 100
	done3 := make(chan struct{})

	randomIntegersStream := streams.RepeatFunctionGenerator[int](done3, func() int {
		return rand.Int()
	})

	time.AfterFunc(delay2, func() {
		close(done3)
	})

	fmt.Printf("\nRandom integers generated during %s:\n", delay2)
	for randomInteger := range randomIntegersStream {
		fmt.Println(randomInteger)
	}

	const (
		numberOfPrimes int = 200
		upperBound     int = 50_000_000
	)

	numberOfPrimeFinders := runtime.NumCPU()
	// numberOfPrimeFinders = 1 // For benchmarking purposes

	now := time.Now().UTC()
	primes := getPrimesUsingFanOutFanIn(numberOfPrimes, upperBound, numberOfPrimeFinders)
	timeToGetPrimes := time.Since(now)

	fmt.Printf("\n%d random prime numbers found in %s using %d goroutines\n",
		numberOfPrimes, timeToGetPrimes, numberOfPrimeFinders)

	fmt.Println("\nRandom prime numbers:")
	for i, prime := range primes {
		fmt.Printf("%d", prime)
		if i < len(primes)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println("")
		}
	}
}

func getPrimesUsingFanOutFanIn(
	numberOPrimes,
	upperBound,
	numberOfPrimeFinders int,
) []int {
	primes := make([]int, 0, numberOPrimes)
	done := make(chan struct{})
	defer close(done)

	randNumbers := streams.RepeatFunctionGenerator[int](done, func() int {
		return rand.Intn(upperBound)
	})

	primeFinders := make([]streams.Stream[int], numberOfPrimeFinders)
	for i := 0; i < numberOfPrimeFinders; i++ {
		primeFinders[i] = randNumbers.Filter(done, isNumberPrime)
	}

	for prime := range streams.
		FanIn[int](done, primeFinders...).
		Take(done, numberOPrimes) {

		primes = append(primes, prime)
	}

	return primes
}

func isNumberPrime(number int) bool {
	squareRoot := int(math.Floor(math.Sqrt(float64(number))))

	for i := 2; i <= squareRoot; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
