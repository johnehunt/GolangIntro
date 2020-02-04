package main

import "fmt"

func max(num1, num2 int) {
	fmt.Printf("max(%v, %v)\n", num1, num2)
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	fmt.Println("result:", result)
}

func main() {
	fmt.Println("Starting")

	// defers the execution of a
	// function until the surrounding function
	// returns (in this case until after mian() completes
	defer max(2, 3)

	fmt.Println("Done")
}
