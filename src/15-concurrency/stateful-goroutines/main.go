package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type ReadOperation struct {
	response chan uint64
}
type WriteOperation struct {
	value    uint64
	response chan bool
}

func statefulWorker(readChannel chan ReadOperation, writeChannel chan WriteOperation) {
	var state uint64 = 0
	for {
		select {
		case read := <-readChannel:
			fmt.Println("Received read operation: ")
			read.response <- state
		case write := <-writeChannel:
			fmt.Println("Received write operation: ", write.value)
			state += write.value
			write.response <- true
		}
	}
}

func reader(readChannel chan ReadOperation, readValue *uint64) {
	fmt.Println("Requesting read operation")
	var responseChannel = make(chan uint64)
	read := ReadOperation{responseChannel}
	readChannel <- read
	r := <-read.response
	fmt.Println("Read response value:", r)
	atomic.AddUint64(readValue, r)
}

func writer(writeChannel chan WriteOperation, writevalue *uint64) {
	var responseChannel = make(chan bool)
	for i := 0; i < 5; i++ {
		fmt.Println("Sending a write operation")
		write := WriteOperation{rand.Uint64(), responseChannel}
		writeChannel <- write
		<-write.response
		atomic.AddUint64(writevalue, 1)
	}
}

func main() {

	var readValue uint64
	var writeValue uint64

	readChannel := make(chan ReadOperation)
	writeChannel := make(chan WriteOperation)

	// state will be owned by a single goroutine
	// In order to read or write that state, other
	// goroutines will send messages to the owning goroutine
	// and receive corresponding replies. These readOp and writeOp
	// structs encapsulate those requests and a way for the
	// owning goroutine to respond.
	go statefulWorker(readChannel, writeChannel)

	// Run a set of Goroutines to read info
	for r := 0; r < 5; r++ {
		go reader(readChannel, &readValue)
	}

	// Run writer Goroutine
	go writer(writeChannel, &writeValue)

	time.Sleep(time.Second)

	finalReadValue := atomic.LoadUint64(&readValue)
	fmt.Println("finalReadValue:", finalReadValue)
	finalWriteValue := atomic.LoadUint64(&writeValue)
	fmt.Println("finalWriteValue:", finalWriteValue)
}
