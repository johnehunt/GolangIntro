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

	printer("A")

	// Invoke function as a goroutine
	// A goroutine is a lightweight thread of execution
	go printer("B")

	fmt.Println("\nX")

	// Pause to give goroutines time to complete
	time.Sleep(time.Second)
	fmt.Println("\nZ")

	fmt.Println("\nDone")
}
