package main

import (
	"fmt"
	"time"
)

func generator(msg string, channel chan string) {
	period := 2
	fmt.Printf("generator > %s Sleep for %d seconds\n", msg, period)
	time.Sleep(time.Duration(period) * time.Second)
	fmt.Printf("generator > sending message\n")
	channel <- msg
}

func worker(channel chan string, done chan bool) {
	for {
		// 2-value form of receive,
		// the more value will be false if jobs has been closed
		// and all values in the channel have already been received
		s, more := <-channel
		if more {
			fmt.Println("worker > Received Task", s)
		} else {
			fmt.Println("worker > Finished Tasks")
			done <- true
			return
		}
	}
}

func main() {
	// Set up two channels
	tasks := make(chan string, 5)
	done := make(chan bool)

	fmt.Println("Start goroutines")
	go worker(tasks, done)
	go generator("A", tasks)

	fmt.Println("Sleep for a bit")
	time.Sleep(3 * time.Second)

	fmt.Println("Close task channel")
	close(tasks)

	fmt.Println("Now await the worker completing")
	<-done

	fmt.Println("Done")
}
