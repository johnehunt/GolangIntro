package main

import (
	"errors"
	"fmt"
)

// In Go it’s idiomatic to communicate errors via an explicit,
// separate return value. T
func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("No division by Zero")
	}
	return x / y, nil
}

func main() {
	fmt.Println("Starting")

	result, err := divide(3, 0)
	if err != nil {
		fmt.Println("There was an error: ", err)
	} else {
		fmt.Println("All was ok:", result)
	}

	fmt.Println("--------------")

	result, err = divide(6, 2)
	if err != nil {
		fmt.Println("There was an error: ", err)
	} else {
		fmt.Println("All was ok:", result)
	}

	fmt.Println("Done")
}
