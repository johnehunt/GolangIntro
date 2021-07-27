package main

import "fmt"

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intChannel := make(chan int)
	go func() {
		defer close(intChannel)
		for _, i := range integers {
			select {
			// If the done channel has been signalled then terminate the goroutine
			case <-done:
				return
			case intChannel <- i:
			}
		}
	}()
	return intChannel
}

func doubler(
	done <-chan interface{},
	intChannel <-chan int,
) <-chan int {
	multipliedIntChannel := make(chan int)
	go func() {
		defer close(multipliedIntChannel)
		for i := range intChannel {
			select {
			case <-done:
				return
			case multipliedIntChannel <- i * 2:
			}
		}
	}()
	return multipliedIntChannel
}

func incrementer(
	done <-chan interface{},
	intChannel <-chan int,
) <-chan int {
	addedIntChannel := make(chan int)
	go func() {
		defer close(addedIntChannel)
		for i := range intChannel {
			select {
			case <-done: // If the done channel has been signalled then terminate the goroutine
				return
			case addedIntChannel <- i + 1:
			}
		}
	}()
	return addedIntChannel
}

func tripler(
	done <-chan interface{},
	intChannel <-chan int,
) <-chan int {
	triplerChannel := make(chan int)
	go func() {
		defer close(triplerChannel)
		for i := range intChannel {
			select {
			case <-done:
				return
			case triplerChannel <- i * 3:
			}
		}
	}()
	return triplerChannel
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
	pipeline := doubler(done, incrementer(done, tripler(done, sourceStream)))

	// Processing data in pipeline
	for v := range pipeline {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
