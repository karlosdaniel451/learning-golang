package main

const maxWorkers int = 32

func handleRPC(RPCQueue chan *RPC) {
	for rpc := range RPCQueue {
		rpc.responseChannel <- rpc.f(rpc.args)
	}
}

func serve(RPCQueue chan *RPC, exitSignal chan struct{}) {
	for i := 0; i < maxWorkers; i++ {
		go handleRPC(RPCQueue)
	}

	<-exitSignal
}

func main() {
	rpcQueue := make(chan *RPC, maxWorkers * 100)
	exitSignal := make(chan struct{})

	serve(rpcQueue, exitSignal)
}
