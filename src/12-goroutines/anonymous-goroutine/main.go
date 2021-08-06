package main

import (
	"fmt"
	"time"
)

func printer(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Print(msg)
	}
}

func main() {
	fmt.Printf("Starting\n\n")

	printer("A") // Run sequentially

	// Invoke function as a goroutine
	// A goroutine is a lightweight thread of execution
	go printer("B")

	fmt.Println("\nX")

	// Can also start a goroutine for an anonymous
	// function call
	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Print(msg)
		}
	}("C")

	fmt.Println("\nY")

	// Pause to give goroutines time to complete
	time.Sleep(time.Second)
	fmt.Println("\nZ")

	fmt.Println("\nDone")
}
