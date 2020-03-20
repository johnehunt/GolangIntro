package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(msg string, channel chan string) {
	period := rand.Int63n(5) + 2
	fmt.Printf("Worker %s Sleep for %d seconds\n", msg, period)
	time.Sleep(time.Duration(period) * time.Second)
	channel <- msg
}

func main() {
	fmt.Println("Starting")
	// Create two channels to select across
	c1 := make(chan string)
	c2 := make(chan string)

	fmt.Println("Launch the goroutines")
	go worker("A", c1)
	go worker("B", c2)

	fmt.Println("Await the messages")
	// use select to block on both channels – will unblock
	// when emssage received on one of the two channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("Done")
}
