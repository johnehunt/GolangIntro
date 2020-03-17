package main

import (
	"container/ring"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Starting")
	// Create a new ring of size 5
	ring := ring.New(5)

	// Get the length of the ring
	ringLength := ring.Len()

	// Initialize the ring with some values
	for i := 0; i < ringLength; i++ {
		ring.Value = rand.Intn(100) + i
		ring = ring.Next()
	}

	// Iterate through the ring and print its contents
	// Higher order function Do takes a functionm to apply
	// to each element in the ring
	ring.Do(func(p interface{}) {
		fmt.Print(p.(int), ",")
	})
	fmt.Println("\nDone")

}
