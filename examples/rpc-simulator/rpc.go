package main

type RPC struct {
	args            []int
	f               func([]int) int
	responseChannel chan<- int
}
