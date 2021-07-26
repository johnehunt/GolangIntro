package main

import "fmt"

func main() {
	// Good programming practice to define constants in CAPITALS

	// Short hand form where type (string) inferred
	const TITLE = "Welcome"
	fmt.Printf("The title is: %s\n", TITLE)

	// Explicit declaration of type for constants
	const LENGTH int = 10
	const WIDTH int = 5

	// Declares area variable of type int with 0 as value
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	// Rate with type float
	const RATE = 3.15
	fmt.Printf("savings rate is: %f\n", RATE)
}
