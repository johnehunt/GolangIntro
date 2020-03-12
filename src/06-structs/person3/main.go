package main

import "fmt"

// Person struct has fields name and age
type Person struct {
	name string
	age  int
}

// NewPerson constructs a new person struct
// with given values. It’s idiomatic to encapsulate
// new struct creation in constructor functions
func newPerson(name string, age int) *Person {
	p := Person{name, age}
	return &p
}

func main() {
	var p = newPerson("John", 21)
	fmt.Println(p)
}
