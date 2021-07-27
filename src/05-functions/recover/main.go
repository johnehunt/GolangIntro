package main

import "fmt"

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

/*
PrintName will generate a Panic if first or last name is an empty string
*/
func PrintName(firstName, lastName string) {
	if firstName == "" {
		panic("error: first name cannot be empty")
	} else if lastName == "" {
		panic("error: last name cannot be empty")
	} else {
		fmt.Println(firstName, lastName)
	}
}

func main() {
	fmt.Println("Starting")
	defer recoverName()

	PrintName("", "Jones")

	fmt.Println("Done")
}
