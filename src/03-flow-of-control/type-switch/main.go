package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	var x interface{}

	// Type switch statement - used to compare types in
	// an emtpy interface
	switch x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case func(int) float64:
		fmt.Println("x is func(int)")
	case bool, string:
		fmt.Println("x is bool or string")
	default: /* Optional */
		fmt.Println("don't know the type")
	}

	fmt.Println("Done")

}
