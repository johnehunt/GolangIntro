package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting")
	// Create a new channel that handles strings
	messageChannel := make(chan string)

	// Use a goroutine to send message
	go func() {
		fmt.Println("Go routine - starting to sleep")
		time.Sleep(2 * time.Second)
		fmt.Println("Go routine - waking up")
		// Send a value into a channel using the send channel <- syntax
		fmt.Println("Go routine - sending message")
		messageChannel <- "Hello World"
		fmt.Println("Go routine - completing")
	}()

	fmt.Println("Waiting for message")
	// The <-channel syntax receives a value from the channel
	msg := <-messageChannel

	// Print out the message received
	fmt.Println(msg)
	fmt.Println("Done")
}
