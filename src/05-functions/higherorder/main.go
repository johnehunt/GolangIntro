package main

import "fmt"

func processor(x int, func1 func(int) int) int {
	return func1(x)
}

func main() {
	fmt.Println("Starting")

	f1 := func(i int) int {
		return i * i
	}

	result := processor(4, f1)
	fmt.Println("result:", result)

	fmt.Println("Done")
}
