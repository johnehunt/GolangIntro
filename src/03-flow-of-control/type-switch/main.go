package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// An idiom used to indicate x is any type of thing
	var x interface{} = "John"

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
	case bool:
		fmt.Println("x is bool")
	case string:
		fmt.Println("x is string")
	default: /* Optional */
		fmt.Println("don't know the type")
	}

	fmt.Println("Done")

}
