package main

import (
	"fmt"
	"time"
)

// Prints a message 10 times
// The done channel is used to notify others
// that the go routie has finished processing
func printer(msg string, done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Print(msg, ", ")
	}
	// Send notification that printer has finished
	done <- true
}

func finalizer(msg string, done chan bool) {
	fmt.Println("Finalizer - Wait for notification")
	// Block until receive notification from
	// the printer goroutine on the channel.
	<-done
	fmt.Println("\nFinalizer - ", msg)
}

func main() {
	fmt.Println("main - Starting")
	// Set up the channel to use for synchronization
	done := make(chan bool, 1)

	// Start up the printer goroutine
	go printer("A", done)
	go finalizer("All Done", done)

	// Pause to give goroutines time to complete
	time.Sleep(time.Second)
	fmt.Println("main - Done")
}
