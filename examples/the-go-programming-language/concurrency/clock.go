package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	var numberOfOpenConnections int

	log.Printf("Waiting for TCP connection requests at %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("Connection accepted with process at %s\n", conn.RemoteAddr())

		wg.Add(1)
		numberOfOpenConnections++

		// This single `go` keyword makes the server able to handle connections
		// concurrently.
		go sendTimeToClient(conn, &wg, &numberOfOpenConnections)
	}
}

func sendTimeToClient(conn net.Conn, wg *sync.WaitGroup, numberOfOpenConnections *int) {
	defer func() {
		conn.Close()
		wg.Done()
		*numberOfOpenConnections--
	}()
	for {
		now := time.Now()
		_, err := io.WriteString(
			conn,
			fmt.Sprintf(
				"%s - number of open connections: %d\n",
				now.Format("15:04:05"), *numberOfOpenConnections,
			),
		)
		if err != nil {
			log.Printf("Connection finished with process at %s\n", conn.RemoteAddr())
			return
		}
		time.Sleep(1 * time.Second)
	}
}
