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

const MAX_GOROUTINES = 5

func main() {
	tasks := make(chan string)

	fmt.Println("Start goroutines")
	for i := 0; i < MAX_GOROUTINES; i++ {
		go worker(tasks)
	}

	tasks <- "Do Work1"
	tasks <- "Do Work2"
	tasks <- "Do Work3"

	fmt.Println("Sleep for a bit")
	time.Sleep(time.Second)
	fmt.Println("Close task channel")
	close(tasks)
	fmt.Println("Sleep for a bit more")
	time.Sleep(time.Second)

	fmt.Println("Done")
}
