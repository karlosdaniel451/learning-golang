package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

type MemBufPool struct {
	pool         *sync.Pool
	bufSize      int
	poolCapacity int
}

func newMemBufPool(bufSize, poolCapacity int) *MemBufPool {
	return &MemBufPool{
		pool: &sync.Pool{
			New: func() any {
				memBuf := make([]byte, bufSize)
				return &memBuf
			},
		},
		bufSize:      bufSize,
		poolCapacity: poolCapacity,
	}
}

func (memBufPool *MemBufPool) getMemBuf() *[]byte {
	return memBufPool.pool.Get().(*[]byte)
}

func (memBufPool *MemBufPool) putMemBuf(memBuf *[]byte) {
	memBufPool.pool.Put(memBuf)
}

func seedMemBuffPool(memBufPool *MemBufPool) {
	for i := 0; i < memBufPool.poolCapacity; i++ {
		memBufPool.pool.Put(memBufPool.pool.New())
	}
}

func intPow(x, y int) int {
	if y == 0 {
		return 1
	}
	return x << (y - 1)
}

func getRandomByte() byte {
	return byte(rand.Intn(255))
}

func createLogFile() *os.File {
	// outputFile, err := os.Open(outputFilename)
	now := time.Now().UTC()
	// outputFile, err := os.OpenFile(
	// 	now.String()+".log", os.O_RDWR|os.O_CREATE,
	// 	fs.ModeAppend,
	// )
	outputFile, err := os.Create(now.String() + ".log")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	return outputFile
}

func main() {
	logFile := createLogFile()
	defer logFile.Close()

	log.SetOutput(logFile)

	fmt.Printf("log file name: %v\n", logFile.Name())

	// Create a MemBuffPool of 1 MB in total, with 1024 buffers of
	// 1 KB each.
	memBufPool := newMemBufPool(intPow(2, 10), intPow(2, 10))
	seedMemBuffPool(memBufPool)

	fmt.Printf("memBufPool.bufSize: %v\n", memBufPool.bufSize)
	fmt.Printf("memBufPool.poolCapacity: %v\n", memBufPool.poolCapacity)

	numberOfWorkers := runtime.NumCPU() * 10

	fmt.Printf("number of logical CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("numberOfWorkers: %v\n", numberOfWorkers)

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			defer wg.Done()

			memBuf := memBufPool.getMemBuf()
			defer memBufPool.putMemBuf(memBuf)

			for i := range *memBuf {
				// (*memBuf)[i] = byte(i)
				(*memBuf)[i] = getRandomByte()
			}

			for i, byteValue := range *memBuf {
				log.Printf("byte written at %d: %08b", i, byteValue)
			}
		}()
	}

	wg.Wait()

	memBuf := memBufPool.getMemBuf()
	// defer memBufPool.putMemBuf(memBuf)

	log.Println("Final memory contents: ")

	for i, byteValue := range *memBuf {
		// log.Printf("%d: %08b", i, byteValue)
		fmt.Fprintf(logFile, "%d: %08b\n", i, byteValue)
	}

}
