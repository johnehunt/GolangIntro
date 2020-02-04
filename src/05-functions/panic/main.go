package main

import "fmt"

func printName(firstName, lastName string) {
	if firstName == "" {
		panic("runtime error: first name cannot be empty")
	} else if lastName == "" {
		panic("runtime error: last name cannot be empty")
	} else {
		fmt.Println(firstName, lastName)
	}
}

func main() {
	fmt.Println("Starting")

	printName("", "Jones")

	fmt.Println("Done")
}
