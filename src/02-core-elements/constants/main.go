package main

import "fmt"

func main() {
	// Good programming practice to define constants in CAPITALS

	const TITLE = "Welcome"
	fmt.Printf("The title is: %s\n", TITLE)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	const RATE = 3.15
	fmt.Printf("savings rate is: %f\n", RATE)
}
