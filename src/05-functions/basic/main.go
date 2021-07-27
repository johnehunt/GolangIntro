package main

import "fmt"

// PrintResult Function without a return result
// Convention is that public functions start with a capital letter
// local functions have an initial lowercase letter
// Public functions should have a comment starting with the name of the function
func PrintResult(i, j int) {
	result := i + j
	fmt.Printf("%d + %d = %d\n", i, j, result)
}

// PrintMsg prints a welcome message
func PrintMsg() {
	fmt.Println("Welcome")
}

// PrintMsg2 prints message based on parameter
func PrintMsg2(msg string) {
	fmt.Println(msg)
}

// Greeter prints a string with name and message
func Greeter(name string, message string) {
	fmt.Printf("Welcome %s - %s\n", name, message)
}

// Max function returns the max value between two integers
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

// inverse illustrates use of named return value and type
// return without a value will implicitly return reciprocal local variable
// can still return other values e.g. return 42
func inverse(v float32) (reciprocal float32) {
	if v == 0 {
		return
	}
	reciprocal = 1 / v
	return
}

func main() {
	fmt.Println("Starting")

	PrintMsg()

	PrintMsg2("Hello World")
	PrintMsg2("Welcome")
	PrintMsg2("Ola")

	Greeter("Eloise", "Hope you like Rugby")

	PrintResult(2, 3)

	fmt.Println("Max(2, 3):", Max(2, 3))

	fmt.Printf("inverse(5)=%.2f\n", inverse(5))

	fmt.Println("Done")
}
