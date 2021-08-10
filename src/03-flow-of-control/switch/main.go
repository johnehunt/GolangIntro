package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	grade := ""
	marks := 90

	// Expression switch statement
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("The mark %d gives you a %s\n", marks, grade)

	// Switch statement without an
	// optional statement
	switch {
	case grade == "A":
		fmt.Println("Excellent!")
	case grade == "B", grade == "C":
		fmt.Println("Well done")
	case grade == "D":
		fmt.Println("You passed")
	case grade == "F":
		fmt.Println("Better try again")
	case marks == 1:
		fmt.Println("Do something about it")
	default:
		fmt.Println("Invalid grade")
	}
	fmt.Printf("Your grade is  %s\n", grade)

	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	var a = 2
	var b = 2

	switch a {
	case b:
		fmt.Println("A same as b")
	case 1:
		fmt.Println("A is one")
	default:
		fmt.Println("Its something else")
	}

	fmt.Println("Done")

}
