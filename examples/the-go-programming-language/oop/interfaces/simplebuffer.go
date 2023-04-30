package main

import (
	"fmt"
	"io"
)

type SimpleBuffer struct {
	content []byte
}

func NewSimpleBuffer(initialLength int) *SimpleBuffer {
	return &SimpleBuffer{content: make([]byte, initialLength)}
}

func (simpleBuffer *SimpleBuffer) Read(p []byte) (n int, err error) {
	if len(simpleBuffer.content) == 0 {
		return 0, io.EOF
	}

	n = copy(p, simpleBuffer.content)
	return n, nil
}

func (simpleBuffer *SimpleBuffer) Write(p []byte) (n int, err error) {
	n = copy(simpleBuffer.content, p)
	return n, nil
}

func (simpleBuffer *SimpleBuffer) String() string {
	return string(simpleBuffer.content)
}

func main() {
	// Assert at compile time that a value of `*SimpleBuffer`` satisfies `io.ReadWriter`.
	var _ io.ReadWriter = new(SimpleBuffer)

	const initialLength int = 100

	simpleBuffer := NewSimpleBuffer(initialLength)
	fmt.Printf("%s\n", simpleBuffer)

	simpleBuffer.Write([]byte("First content"))
	fmt.Printf("%s\n", simpleBuffer)

	simpleBuffer.Write([]byte("Second content"))
	fmt.Printf("%s\n", simpleBuffer)

	currentContent := make([]byte, initialLength)
	currentLength, _ := simpleBuffer.Read(currentContent)
	fmt.Printf("%s\n", currentContent[:currentLength])

	simpleBuffer.Write([]byte("Third content"))
	currentContent = make([]byte, initialLength)
	currentLength, _ = simpleBuffer.Read(currentContent)
	fmt.Printf("%s\n", currentContent[:currentLength])
}
