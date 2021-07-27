package main

import (
	"fmt"
	"time"
)

func worker(channel chan string) {
	for {
		// 2-value form of receive,
		// the more value will be false if jobs has been closed
		// and all values in the channel have already been received
		s, more := <-channel
		if more {
			fmt.Println("worker > Received Task", s)
		} else {
			fmt.Println("worker > Finished Tasks")
			return
		}
	}
}

func main() {
	// Set up the channel
	tasks := make(chan string, 5)

	fmt.Println("Start goroutines")
	go worker(tasks)
	tasks <- "Do Work1"
	tasks <- "Do Work2"

	fmt.Println("Sleep for a bit")
	time.Sleep(3 * time.Second)

	fmt.Println("Close task channel")
	close(tasks)

	fmt.Println("Done")
}
