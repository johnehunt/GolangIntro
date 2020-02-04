package main

import "fmt"

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func printName(firstName, lastName string) {
	defer recoverName()
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
