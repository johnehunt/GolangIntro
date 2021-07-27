package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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

func main() {

	fmt.Println("Starting - setting up data")
	size := 10000

	data := make([]int, size, size)
	for n := 0; n < size; n++ {
		data = append(data, rand.Intn(size))
	}

	fmt.Println("Starting calculations for Prime numbers")
	start := time.Now()

	for _, n := range data {
		fmt.Printf("%d is prime: %v\n", n, isPrime(n))
	}

	fmt.Printf("Search took: %v\n", time.Since(start))

}
