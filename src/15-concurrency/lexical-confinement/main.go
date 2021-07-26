package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// Channel is created within the lexical scope of
	// the function. As the returned channel is read only channel
	// anything external to the function can only read data
	// form the channel. Internally the goroutine can write data to
	// the channel.
	channelOwner := func() <-chan int {
		results := make(chan int, 5)
		// hidden write orieinted goroutine
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// This function is passed a read only channel and thus
	// it is not possible for the function to 'accidentally' write
	// to the channel - it can therefore only consume data
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := channelOwner()
	consumer(results)

	fmt.Println("Done")
}
