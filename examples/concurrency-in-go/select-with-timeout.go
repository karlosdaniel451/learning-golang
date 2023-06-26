package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	outputStream := streamRead(os.Stdin)
	numberOfReads := 0

	const maxIdleTimeInSecs = time.Second * 5
	fmt.Printf("maxIdleTime: %.3f secs\n\n", maxIdleTimeInSecs.Seconds())

	loopCond := true
	for loopCond {
		select {
		case output := <-outputStream:
			numberOfReads++
			fmt.Printf("%s\n\n", output)
		case <-time.After(maxIdleTimeInSecs):
			fmt.Printf("\nTimed out after %.3f secs idle!\n", maxIdleTimeInSecs.Seconds())
			loopCond = false
		}
	}

	fmt.Printf("numberOfReads: %d\n", numberOfReads)
}

func streamRead(ioReader io.Reader) (outputStream chan []byte) {
	outputStream = make(chan []byte)

	scanner := bufio.NewScanner(ioReader)
	go func() {
		defer close(outputStream)
		for scanner.Scan() {
			outputStream <- scanner.Bytes()
		}
	}()

	return outputStream
}
