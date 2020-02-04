package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	x := 20
	y := 10

	// Simple if ... statement
	if x > y {
		/* if condition is true then print the following */
		fmt.Println("X is bigger")
	}

	// If ... else ... statement
	if x < y {
		/* if condition is true then print the following */
		fmt.Println("x is less than y")
	} else {
		/* if condition is false then print the following */
		fmt.Println("x is not less than y")
	}

	// if ... else if ... else ... statement
	if x == 10 {
		/* if condition is true then print the following */
		fmt.Println("Value of x is 10")
	} else if x == 20 {
		/* if else if condition is true */
		fmt.Println("Value of x is 20")
	} else if x == 30 {
		/* if else if condition is true  */
		fmt.Println("Value of x is 30")
	} else {
		/* if none of the conditions is true */
		fmt.Println("None of the values is matching")
	}

	// Can nest If statements

	if x == 20 {
		/* if condition is true then check the following */
		if y == 10 {
			/* if condition is true then print the following */
			fmt.Println("Value of x is 20 and y is 10")
		}
	}

	fmt.Println("Done")

}
