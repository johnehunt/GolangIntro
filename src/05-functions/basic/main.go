package main

import "fmt"

// PrintResult Functionw without a return result
// Convention is the public functions start with a capital letter
// local functions have an initial lowercase letter
// Public functions shoudl have a comment starting with the name of the function
func PrintResult(i int, j int) {
	result := i + j
	fmt.Printf("%d + %d = %d\n", i, j, result)
}

// Max function returning the max between two numbers
func Max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	fmt.Println("Starting")

	PrintResult(2, 3)

	fmt.Println("Max(2, 3):", Max(2, 3))

	fmt.Println("Done")
}
