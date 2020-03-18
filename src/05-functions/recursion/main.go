package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1 // Termination condition
	}
	return i * factorial(i-1) // Recursive call
}

func main() {
	fmt.Println("Starting")

	var i int = 15
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	fmt.Println("Done")
}
