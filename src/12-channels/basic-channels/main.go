package main

import "fmt"

func main() {

	// Create a new channel that handles strings
	messages := make(chan string)

	// Use a goroutine to send message
	go func() {
		// Send a value into a channel using the channel <- syntax
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel
	msg := <-messages

	// Print out the message received
	fmt.Println(msg)
}
