package main

import "fmt"

// Max compares two number and returns max
func Max(num1, num2 int) {
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
	defer Max(2, 3)

	fmt.Println("Done")
}
