package main

import (
	"fmt"
	"os"
)

func main() {
	var i, j int // ints intialized to zero
	fmt.Printf("Values of i: %d and j: %d\n", i, j)

	var flag, x, total = true, 23, 42.5
	fmt.Printf("Values of flag: %t, x: %d and total: %f", flag, x, total)

	// Funcrtions can return multiple values
	// variables can be set based on each
	var file, err = os.Open("temp.txt")
	fmt.Println(file)
	fmt.Println(err)

}
