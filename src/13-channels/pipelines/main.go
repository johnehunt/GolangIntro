package main

import "fmt"

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			// If the done channel has been signalled then terminate the goroutine
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(
	done <-chan interface{},
	intStream <-chan int,
	multiplier int,
) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func add(
	done <-chan interface{},
	intStream <-chan int,
	additive int,
) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done: // If the done channel has been signalled then terminate the goroutine
				return
			case addedStream <- i + additive:
			}
		}
	}()
	return addedStream
}

func main() {
	fmt.Println("Starting")

	// create the done channel and ensure that done is closed
	// when the program terminates
	done := make(chan interface{})
	defer close(done)

	// Set up the data generator
	sourceStream := generator(done, 1, 2, 3, 4)

	// Set up the pipeline
	pipeline := multiply(done, add(done, multiply(done, sourceStream, 2), 1), 2)

	// Processing data in pipeline
	for v := range pipeline {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
