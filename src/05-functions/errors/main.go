package main

import (
	"errors"
	"fmt"
)

func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("Can't divide by Zero")
	}
	return x / y, nil
}

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func divide2(x int, y int) (int, error) {
	if y == 0 {
		return -1, &argError{y, "Can't divide by Zero"}
	}
	return x / y, nil
}

func main() {
	fmt.Println("Starting")

	result, err := divide(6, 0)
	if err != nil {
		fmt.Println("There was a problem:", err)
	} else {
		fmt.Println("result:", result)
	}

	// Short hand form for if
	if result, err = divide2(6, 0); err != nil {
		fmt.Println("There was a problem:", err)
	} else {
		fmt.Println("result:", result)
	}

	fmt.Println("Done")
}
