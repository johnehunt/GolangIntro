package main

import "fmt"

/*
PrintName illustrates how panic can be used.
Note this is not considered the way in whioch functiosn shoudl indicate standard errors in Go
Panic shoudl only be used in very exceptional situaitons
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

	PrintName("", "Jones")

	fmt.Println("Done")
}
