package main

import "fmt"

// type RPCArg interface{}
// type RPCReturnValue interface{}

// type RPC struct {
// 	args []RPCArg
// 	// returnValue ProcedureCallReturnValue
// 	f func([]RPCArg) RPCReturnValue
// 	resultchan RPCReturnValue
// }

func sum(integers []int) int {
	total := 0
	for _, value := range integers {
		total += value
	}
	return total
}

func main() {
	rpcChannel := make(chan RPC)

	rpc1 := RPC{[]int{1, 2, 3}, sum, make(chan int)}
	rpcChannel <- rpc1

	rpc2 := RPC{[]int{1, 2, 3, 4}, sum, make(chan int)}
	rpcChannel <- rpc2

	fmt.Printf("1st Return value: %d\n", rpc1.responseChannel)
	fmt.Printf("2st Return value: %d\n", rpc2.responseChannel)
}
