package main

import "fmt"

func main() {
	fmt.Println("Starting")

	// Function literal stored into variable func1
	var func1 = func(i int) int {
		return i * i
	}

	fmt.Printf("func1: %v\n", func1)
	fmt.Printf("func1(4): %d\n", func1(4))

	// function variables are still strongly typed
	var func2 func(int) int = func1
	fmt.Printf("func2(4): %d\n", func2(4))

	// creating and invoking an anonymous function in one go
	var x = func() int {
		return 2 + 3
	}()

	fmt.Println("Value of x is:", x)

	// function without a return type
	var func3 func(int, int) = func(x int, y int) {
		fmt.Println(x + y)
	}
	func3(3, 4)

	fmt.Println("Done")
}
