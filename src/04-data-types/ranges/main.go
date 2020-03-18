package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// prints contents of numbers
	numbers := []int{0, 1, 7, 3}
	for i, e := range numbers {
		fmt.Printf("Item @ %d = %d\n", i, e)
	}

	// print elements of a string
	for s := range "John" {
		fmt.Printf("s: %v; ", s)
	}
	fmt.Println()

	fmt.Println("Done")
}
