package main

import "fmt"

func main() {

	// Create a new channel that handles strings
	messageChannel := make(chan string)

	// Use a goroutine to send message
	go func() {
		// Send a value into a channel using the channel <- syntax
		messageChannel <- "Hello World"
	}()

	// The <-channel syntax receives a value from the channel
	msg := <-messageChannel

	// Print out the message received
	fmt.Println(msg)
}
