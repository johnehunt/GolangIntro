package main

import "fmt"

// Functionw without a return result
func printResult(i int, j int) {
	result := i + j
	fmt.Printf("%d + %d = %d\n", i, j, result)
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
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

	printResult(2, 3)

	fmt.Println("max(2, 3):", max(2, 3))

	fmt.Println("Done")
}
