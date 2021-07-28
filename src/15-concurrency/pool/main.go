package main

import (
	"fmt"
	"sync"
)

type Info struct {
	ID   int
	Data string
}

func (i *Info) Reset() {
	i.ID = 0
	i.Data = ""
}

func main() {
	fmt.Println("Starting")

	// Set up Pool
	var infoPool = sync.Pool{
		// New creates an instance when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			return new(Info)
		},
	}

	// Obtain an object from the Pool
	// note need to cast it to a pointer to Info
	info := infoPool.Get().(*Info)
	info.ID = 23
	info.Data = "generla info"
	fmt.Println(*info)

	// Once finished with it - reset it and return to Pool
	info.Reset()
	infoPool.Put(info)

	fmt.Println("Done")
}
