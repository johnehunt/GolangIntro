package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Provides a simple fan in - fan out example

func generator(numbers []int) <-chan int {
	intStream := make(chan int)
	go func() {
		for _, n := range numbers {
			intStream <- n
		}
		close(intStream)
	}()
	return intStream
}

func isPrime(value int) bool {
	// Zero and 1 are not considered prime numbers
	if value == 0 || value == 1 {
		return false
	}
	// The only even prime number is 2.
	if value == 2 {
		return true
	}
	// For any other number determien if its a prime
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func primeNumberCheck(intChannel <-chan int) <-chan bool {
	resultsChannel := make(chan bool)
	go func() {
		for i := range intChannel {
			resultsChannel <- isPrime(i)
		}
		close(resultsChannel)
	}()
	return resultsChannel
}

func fanIn(c1, c2, c3, c4 <-chan bool) <-chan bool {
	combinedResultsChannel := make(chan bool)
	go func() {
		for {
			select {
			case s := <-c1:
				combinedResultsChannel <- s
			case s := <-c2:
				combinedResultsChannel <- s
			case s := <-c3:
				combinedResultsChannel <- s
			case s := <-c4:
				combinedResultsChannel <- s
			}
		}
	}()
	return combinedResultsChannel
}

func main() {
	fmt.Println("Starting - setting up data")
	const SIZE = 10000

	data := make([]int, SIZE, SIZE)
	for n := 0; n < SIZE; n++ {
		data = append(data, rand.Intn(SIZE))
	}

	// Create shared input channel
	inputChannel := generator(data)

	// Fan out to 4 goroutines
	c1 := primeNumberCheck(inputChannel)
	c2 := primeNumberCheck(inputChannel)
	c3 := primeNumberCheck(inputChannel)
	c4 := primeNumberCheck(inputChannel)

	// Now fan in and combine results
	results := fanIn(c1, c2, c3, c4)

	for i := 0; i < len(data); i++ {
		fmt.Printf("%d is prime: %v\n", data[i], <-results)
	}
}
