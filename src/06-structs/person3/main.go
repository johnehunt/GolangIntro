package main

import "fmt"

// Person struct has fields Name and Age
type Person struct {
	Name string
	Age  int
}

// NewPerson constructs a new person struct
// with given values. It’s idiomatic to encapsulate
// new struct creation in constructor functions
func newPerson(Name string, Age int) *Person {
	p := Person{Name, Age}
	return &p
	// alternative use - new keyword can be used
	// to create a new struct. It returns a pointer
	// to the newly created struct
	// p := new(Person)
	// p.Name = Name
	// p.Age = Age
	// return p
}

func main() {
	var p = newPerson("John", 21)
	fmt.Println(p)
}
