package main

import "fmt"

// Processor applies a function to an int
func Processor(x int, func1 func(int) int) int {
	return func1(x)
}

func main() {
	fmt.Println("Starting")

	f1 := func(i int) int {
		return i * i
	}

	result := Processor(4, f1)
	fmt.Println("result:", result)

	fmt.Println("Done")
}
