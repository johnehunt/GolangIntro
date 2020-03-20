package main

import (
	"fmt"
	"time"
)

// MessageData for use with Channels
type MessageData struct {
	ID   string
	info string
}

func printer(channel chan MessageData) {

	for true {
		msg := <-channel
		fmt.Println(msg)
	}

}

func main() {
	fmt.Println("Starting")

	messageDataChannel := make(chan MessageData)

	go printer(messageDataChannel)

	messageDataChannel <- MessageData{"1", "Hello World"}
	messageDataChannel <- MessageData{"2", "Welcome"}
	messageDataChannel <- MessageData{"3", "Ola"}
	messageDataChannel <- MessageData{"4", "Bonjour"}

	// Give time for all messages to be delivered
	time.Sleep(time.Second)

	fmt.Println("Done")
}
