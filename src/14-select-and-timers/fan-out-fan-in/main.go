package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// Provides a simple fan in - fan out example

func generatePipeline(numbers []int) <-chan int {
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

func primeNumberCheckStage(intChannel <-chan int) <-chan bool {
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

	numFinders := runtime.NumCPU()
	fmt.Printf("Number of CPUS %d\n", numFinders)

	fmt.Println("Starting - setting up data")
	size := 10000

	data := make([]int, size, size)
	for n := 0; n < size; n++ {
		data = append(data, rand.Intn(size))
	}

	start := time.Now()

	// Create shared input channel
	inputChannel := generatePipeline(data)

	// Fan out to 4 goroutines
	c1 := primeNumberCheckStage(inputChannel)
	c2 := primeNumberCheckStage(inputChannel)
	c3 := primeNumberCheckStage(inputChannel)
	c4 := primeNumberCheckStage(inputChannel)

	// Now fan in and combine results
	results := fanIn(c1, c2, c3, c4)

	for i := 0; i < len(data); i++ {
		fmt.Printf("%d is prime: %v\n", data[i], <-results)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))

}
