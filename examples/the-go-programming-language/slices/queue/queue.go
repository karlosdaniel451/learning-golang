package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []int
}

func main() {
	queue := createQueue()

	enqueue(&queue, 1)
	enqueue(&queue, 2)
	enqueue(&queue, 3)
	fmt.Println(queue)

	dequeue(&queue)
	fmt.Println(queue)

	dequeue(&queue)
	fmt.Println(queue)

	enqueue(&queue, 4)
	fmt.Println(queue)
}

func createQueue() Queue {
	return Queue{elements: make([]int, 0)}
}

func enqueue(queue *Queue, element int) {
	queue.elements = append(queue.elements, element)
}

func dequeue(queue *Queue) (int, error) {
	if len(queue.elements) == 0 {
		return 0, errors.New("the given queue is empty")
	}

	element := queue.elements[0]

	queue.elements = queue.elements[1:]

	return element, nil
}
