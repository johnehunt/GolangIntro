package main

import (
	"fmt"
)

var name string

// Function called when program is run before the main method
func init() {
	fmt.Println("Init function on main initialization")
	name = "John"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello", name)
}
