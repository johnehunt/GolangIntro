package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// Employee struct
// Using an embedded type (Person) in Employee
// AKA anonymous fields
type Employee struct {
	Person
	id string
}

func main() {
	e := new(Employee)
	e.id = "EMP987"
	// Long hand reference
	e.Person.name = "John"
	// Shorthand form
	e.age = 21
	fmt.Println("e:", e)

	var e1 = Employee{Person: Person{"Denise", 21}, id: "EMP123"}
	fmt.Println("e1:", e1)

	var e2 = Employee{Person{"Denise", 21}, "EMP123"}
	fmt.Println("e2:", e2)

	var e3 = Employee{id: "EMP345"}
	fmt.Println("e3:", e3)

	var e4 = Employee{Person: Person{"Denise", 21}}
	fmt.Println("e4:", e4)

}
