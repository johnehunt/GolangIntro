package main

import (
	"fmt"
	"os"
)

func main() {
	var i, j int // ints intialized to zero
	fmt.Printf("Values of i: %d and j: %d\n", i, j)

	// Can do multiple declarations and initializations in a line
	var flag, x, total = true, 23, 42.5
	fmt.Printf("Values of flag: %t, x: %d and total: %f\n", flag, x, total)

	// Functions can return multiple values
	// variables can be set based on each
	var file, err = os.Open("temp.txt")

	// %v is a catch all for printing a value
	fmt.Printf("file: %v\n", file)
	fmt.Printf("err: %v\n", err)

}
