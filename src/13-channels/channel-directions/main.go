package main

import (
	"fmt"
	"time"
)

// This ping function only accepts a string channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	time.Sleep(time.Second)
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and
// a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	time.Sleep(time.Second)
	pongs <- msg
}

func main() {
	fmt.Println("Starting")

	fmt.Println("Set up channels")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	fmt.Println("Send message")
	ping(pings, "Hello")
	pong(pings, pongs)

	fmt.Println("Wait for response")
	fmt.Println(<-pongs)
	fmt.Println("Done")
}
