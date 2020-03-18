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

	// Slightly extended version
	if x > y {
		fmt.Println("X is bigger")
		fmt.Printf("X is %d, x * x = %d\n", x, x*x)
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

	snowing := true
	temp := -1
	if temp < 0 {
		fmt.Println("It is freezing")
		/* if condition is true then check the following */
		if snowing {
			/* if condition is true then print the following */
			fmt.Println("Put on boots")
		}
		fmt.Println("Time for Hot Chocolate")
	}

	fmt.Println("Done")

	// Can use logical operators
	age := 15
	status := ""
	if age > 12 && age < 20 {
		status = "teenager"
	} else {
		status = "not teenager"
	}
	fmt.Printf("You are a %s\n", status)

}
