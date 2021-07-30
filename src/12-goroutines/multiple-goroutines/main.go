package main

import (
	"fmt"
	"time"
)

func worker(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(msg)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting")

	// fire off set of goroutines
	go worker("A", 10)
	go worker("B", 15)
	go worker("C", 8)
	go worker("D", 5)

	// Let goroutinmes run
	time.Sleep(time.Second)

	fmt.Println("\nDone")
}
