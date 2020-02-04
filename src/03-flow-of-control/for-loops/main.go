package main

import (
	"fmt"
)

func main() {
	// Basic For loop
	for i := 0; i < 5; i++ {
		fmt.Printf("value of i: %v\n", i)
	}

	fmt.Println("------------------")

	// For loop as a while loop
	x := 0
	for x < 5 {
		fmt.Printf("value of x: %v\n", x)
		x++
	}

	fmt.Println("------------------")

	// Nested for
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%v * %v is %v\n", i, j, i*j)
		}
	}

	fmt.Println("------------------")

	// Not using a loop variable
	// Note for i=0 not i := 0
	var i = 0
	for i = 0; i < 4; i++ {
		fmt.Printf("value of i: %v\n", i)
	}
}
