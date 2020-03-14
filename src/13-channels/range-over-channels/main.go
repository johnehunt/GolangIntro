package main

import "fmt"

func main() {
	// iterate over 2 values in the queue channel.

	fmt.Println("Setting up channel")
	queue := make(chan string, 2)
	queue <- "A"
	queue <- "B"
	fmt.Println("Closing channel")
	close(queue)

	// range iterates over each element as it’s
	// received from queue. Because the channel is closed above,
	// the iteration terminates after receiving the 2 elements.
	fmt.Println("Process data in channel")
	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("Done")
}
