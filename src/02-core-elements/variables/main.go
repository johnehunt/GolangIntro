package main

import (
	"fmt"
	"os"
)

func main() {

	var address string
	address = "10 High Street"
	fmt.Println(address)

	var i, j int // ints intialized to zero
	fmt.Printf("Values of i: %d and j: %d\n", i, j)

	var title string = "Dr"
	var surname = "Cooke"
	fmt.Printf("%s %s\n", title, surname)

	// Can do multiple declarations and initializations in a line
	var flag, x, total = true, 23, 42.5
	fmt.Printf("Values of flag: %t, x: %d and total: %4.2f\n", flag, x, total)

	// Can define set fo different variables and types using var (..)
	var (
		name   = "John"
		age    = 57
		height int
	)

	fmt.Printf("%s is %d height: %d\n", name, age, height)

	// Functions can return multiple values
	// variables can be set based on each
	var file, err = os.Open("temp.txt")

	// %v is a catch all for printing a value
	fmt.Printf("file: %v\n", file)
	fmt.Printf("err: %v\n", err)

}
