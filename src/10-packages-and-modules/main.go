package main

import (
	"10-packages-and-modules/packages/calc"
	"fmt"
)

func main() {
	var result, err = calc.Calculator("+", 2, 3)
	fmt.Println("Result:", result)
}
