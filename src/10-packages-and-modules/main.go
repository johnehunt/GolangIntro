package main

import (
	"10-packages-and-modules/modules/calc"
	"fmt"
)

func main() {
	var result, err = calc.Calculator("+", 2, 3)
	if err != nil {
		fmt.Println("Something went wrong", err)
	} else {
		fmt.Println("Result:", result)
	}
}
