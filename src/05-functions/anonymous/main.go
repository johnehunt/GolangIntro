package main

import "fmt"

func main() {
	fmt.Println("Starting")
	var x = func() int {
		return 2 + 3
	}()
	fmt.Println("Value of x is:", x)
	fmt.Println("Done")
}
