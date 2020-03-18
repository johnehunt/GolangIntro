package main

import "fmt"

// Person struct has fields Name and Age
type Person struct {
	Name string
	Age  int
}

func printPerson(person Person) {
	fmt.Println(person)
	fmt.Println(person.Name)
}

func printPerson2(person *Person) {
	fmt.Println(*person)
	// person pointer is automatically de referenced
	fmt.Println(person.Name)
	// manually dereferencing
	fmt.Println((*person).Name)
}

func main() {
	// Creating and initializing a struct
	// with a 'struct literal'
	var p1 = Person{Name: "John", Age: 21}
	printPerson(p1)
	printPerson2(&p1)
}
