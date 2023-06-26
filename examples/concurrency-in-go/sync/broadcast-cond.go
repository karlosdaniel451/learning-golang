package main

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	clickEvent *sync.Cond
	hoverEvent *sync.Cond
}

func onEvent(eventCond *sync.Cond, eventHandler func()) {
	goroutineStarted := make(chan struct{})
	go func() {
		goroutineStarted <- struct{}{}
		eventCond.L.Lock()
		eventCond.Wait()
		eventHandler()
		defer eventCond.L.Unlock()
	}()
	<-goroutineStarted
}

func newButton() *Button {
	return &Button{
		clickEvent: sync.NewCond(&sync.Mutex{}),
		hoverEvent: sync.NewCond(&sync.Mutex{}),
	}
}

func (button *Button) onClick(eventHandler func()) {
	onEvent(button.clickEvent, eventHandler)
}

func (button *Button) onHover(eventHandler func()) {
	onEvent(button.hoverEvent, eventHandler)
}

func clickButton(button *Button) {
	button.clickEvent.Broadcast()
}

func hoverButton(button *Button) {
	button.hoverEvent.Broadcast()
}

func main() {
	var wg sync.WaitGroup

	button := newButton()

	wg.Add(1)
	button.onClick(func() {
		fmt.Println("Button was clicked.")
		wg.Done()
	})

	wg.Add(1)
	button.onClick(func() {
		fmt.Println("Button was clicked 2.")
		wg.Done()
	})

	wg.Add(1)
	button.onHover(func() {
		fmt.Println("Button was hovered.")
		wg.Done()
	})

	time.Sleep(time.Second * 1)
	clickButton(button)

	time.Sleep(time.Second * 1)
	hoverButton(button)

	wg.Wait()
}
