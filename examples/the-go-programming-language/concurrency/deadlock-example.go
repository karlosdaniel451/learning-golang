package main

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func1(ch1, ch2)
	func1(ch2, ch1)
}

func func1(out chan<- struct{}, in <-chan struct{}) {
	out <- struct{}{}
	<-in
}
