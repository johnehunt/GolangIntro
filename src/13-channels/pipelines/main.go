package main

import "fmt"

/* Note returns readonly channel
   internally function can write to channel
   but externally others can only read from it.
   This is called lexical confinement
*/
func generator(integers ...int) <-chan int {
	intChannel := make(chan int)
	go func() {
		for _, i := range integers {
			intChannel <- i
		}
		close(intChannel)
	}()
	return intChannel
}

func doubler(intChannel <-chan int) <-chan int {
	multipliedIntChannel := make(chan int)
	go func() {
		for i := range intChannel {
			multipliedIntChannel <- i * 2
		}
		close(multipliedIntChannel)
	}()
	return multipliedIntChannel
}

func incrementer(intChannel <-chan int) <-chan int {
	addedIntChannel := make(chan int)
	go func() {
		for i := range intChannel {
			addedIntChannel <- i + 1
		}
		close(addedIntChannel)
	}()
	return addedIntChannel
}

func tripler(intChannel <-chan int) <-chan int {
	triplerChannel := make(chan int)
	go func() {
		for i := range intChannel {
			triplerChannel <- i * 3
		}
		close(triplerChannel)
	}()
	return triplerChannel
}

func main() {
	fmt.Println("Starting")

	// Set up the data generator
	sourceStream := generator(1, 2, 3, 4)

	// Set up the pipeline
	pipeline := doubler(incrementer(tripler(sourceStream)))

	// Processing data in pipeline
	for v := range pipeline {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
