package main

import "fmt"

// Person struct implements Stringer interface.
// The Stringer interface - has one method
// String() string.package main
// A Stringer is a type that can describe itself as a string.
type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.name, p.age)
}

func main() {
	a := Person{"John", 21}
	b := Person{"Phoebe", 18}
	fmt.Println(a, ", ", b)
}
