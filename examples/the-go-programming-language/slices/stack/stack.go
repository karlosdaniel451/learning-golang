package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []int
}

func main() {
	stack := createStack()

	push(&stack, 1)
	push(&stack, 2)
	push(&stack, 3)
	fmt.Println(stack)

	pop(&stack)
	fmt.Println(stack)

	pop(&stack)
	fmt.Println(stack)

	push(&stack, 4)
	fmt.Println(stack)
}

func createStack() Stack {
	return Stack{elements: make([]int, 0)}
}

func push(queue *Stack, element int) {
	queue.elements = append([]int{element}, queue.elements...)
}

func pop(queue *Stack) (int, error) {
	if len(queue.elements) == 0 {
		return 0, errors.New("the given stack is empty")
	}

	element := queue.elements[0]

	queue.elements = queue.elements[1:]

	return element, nil
}
