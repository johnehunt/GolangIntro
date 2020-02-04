package main

import "fmt"

func main() {
	fmt.Println("Starting")

	var func1 = func(i int) int {
		return i * i
	}

	fmt.Println("func1:", func1)
	fmt.Println("func1(4):", func1(4))

	var func2 func(int) int = func1
	fmt.Println("func2:", func2)
	fmt.Println("func2(4):", func2(4))

	fmt.Println("Done")
}
