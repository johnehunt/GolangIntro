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
	fmt.Println("Starting")

	printer("A")

	// Invoke function as a goroutine
	// A goroutine is a lightweight thread of execution
	go printer("B")

	fmt.Print("X")

	// Can also start a goroutine for an anonymous
	// function call
	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Print(msg)
		}
	}("C")

	fmt.Print("Y")

	// Pause to give goroutines time to complete
	time.Sleep(time.Second)
	fmt.Print("Z")

	fmt.Println("\nDone")
}
