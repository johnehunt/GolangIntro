package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting")

	// Set up Pool
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	// Obtain an object from the Pool
	buffer := bufferPool.Get().(*bytes.Buffer)

	// Once finished with it - reset it and return to Pool
	buffer.Reset()
	bufferPool.Put(buffer)

	fmt.Println("Done")
}
