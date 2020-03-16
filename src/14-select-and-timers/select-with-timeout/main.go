package main

import (
	"fmt"
	"time"
)

func worker(msg string, channel chan string) {
	period := 2
	fmt.Printf("Worker %s Sleep for %d seconds\n", msg, period)
	time.Sleep(time.Duration(period) * time.Second)
	channel <- msg
}

func main() {
	fmt.Println("Starting")

	fmt.Println("Set up channel and 1st goroutine")
	c1 := make(chan string, 1)
	go worker("A", c1)

	fmt.Println("select with 1 second timeout")
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	fmt.Println("Set up channel and 1st goroutine")
	c2 := make(chan string, 1)
	go worker("B", c2)

	fmt.Println("select with 4 second timeout")
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout 2")
	}
	fmt.Println("Done")
}
