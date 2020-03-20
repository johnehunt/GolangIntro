package main

import "fmt"

func main() {

	// make a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// Send to message and buffer them
	messages <- "Hello"
	messages <- "World"
	// messages <- "Brave New World" // will result in an error

	// Receive the two buffered messages
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
