package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Employee struct
// Using an embedded type (Person) in Employee
// AKA anonymous fields
type Employee struct {
	Person
	ID string
}

func main() {
	e := Employee{}
	e.ID = "EMP987"
	// Long hand reference
	e.Person.Name = "John"
	// Shorthand form
	e.Age = 21
	fmt.Println("e:", e)

	var e1 = Employee{
		Person: Person{"Denise", 21},
		ID:     "EMP123", // Need a comma before a newline in a struct literal
	}
	fmt.Println("e1:", e1)

	var e2 = Employee{
		Person{"Denise", 21},
		"EMP123"}
	fmt.Println("e2:", e2)

	var e3 = Employee{ID: "EMP345"}
	fmt.Println("e3:", e3)

	var e4 = Employee{Person: Person{"Denise", 21}}
	fmt.Println("e4:", e4)

}
