package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func worker(counter *uint64, waitGroup *sync.WaitGroup) {
	for c := 0; c < 1000; c++ {
		// To atomically increment the counter use AddUint64,
		// giving it the memory address of the ops counter
		atomic.AddUint64(counter, 1)
	}
	waitGroup.Done()
}

func main() {
	fmt.Println("Starting")
	var counter uint64

	// WaitGroup will wait for all goroutines to finish
	var waitGroup sync.WaitGroup

	// start 50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		waitGroup.Add(1)
		go worker(&counter, &waitGroup)
	}

	// Wait until all the goroutines are done.
	waitGroup.Wait()

	// Safe to access ops as no other goroutine is now writing
	fmt.Println("counter:", counter)
	fmt.Println("Done")
}
