package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non blocking receive using a default select
	// If a value is available on messages select
	// will handle it with <-message case.
	// If not it will immeidately take the default case
	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	default:
		fmt.Println("No message received")
	}

	// A non-blocking send works similarly.
	// In this case the msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Message sent", msg)
	default:
		fmt.Println("No message sent")
	}

	// can use multiple cases above the default clause to
	// implement a multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	case sig := <-signals:
		fmt.Println("Signal received", sig)
	default:
		fmt.Println("No activity")
	}
}
