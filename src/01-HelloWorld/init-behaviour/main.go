package main

import (
	"fmt"
)

var name string

func init() {
	fmt.Println("Init Method on main initialization")
	name = "John"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello", name)
}
