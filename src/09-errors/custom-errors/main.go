package main

import (
	"fmt"
)

// InvalidArgumentError custom exception type
// Implements the Error interface with the Errors method
type InvalidArgumentError struct {
	Argument int
	Problem  string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.Argument, e.Problem)
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, &InvalidArgumentError{y, "No division by Zero"}
	}
	return x / y, nil
}

func main() {
	fmt.Println("Starting")

	result, err := divide(3, 0)
	if err != nil {
		fmt.Println("There was an error: ", err)
		// Use type assertion to get data from exception
		ae := err.(*InvalidArgumentError)
		if ae != nil {
			fmt.Println(ae.Argument)
			fmt.Println(ae.Problem)
		}
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
