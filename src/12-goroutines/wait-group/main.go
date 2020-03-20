package main

import (
	"fmt"
	"sync"
)

func printer(msg string, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we’re done.
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Print(msg)
	}
}
func main() {
	fmt.Println("Starting")

	// Waitgroup to wait until a group of goroutines finish
	fmt.Println("Setting up the wait group")
	var wg sync.WaitGroup

	fmt.Println("Starting Goroutines")
	// Launch several goroutines and
	// increment the WaitGroup counter for each
	wg.Add(1)
	go printer("A", &wg)
	wg.Add(2)
	go printer("B", &wg)
	go printer("C", &wg)

	// Wait for goroutines to complete
	fmt.Println("Waiting for Goroutines to finish")
	wg.Wait()

	fmt.Println("\nDone")
}
