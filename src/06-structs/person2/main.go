package main

import "fmt"

// Person struct has fields name and age
type Person struct {
	name string
	age  int
}

func printPerson(person Person) {
	fmt.Println(person)
	fmt.Println(person.name)
}

func printPerson2(person *Person) {
	fmt.Println(*person)
	// person pointer is automatically de referenced
	fmt.Println(person.name)
	// manually dereferencing
	fmt.Println((*person).name)
}

func main() {
	var p1 = Person{name: "John", age: 21}
	printPerson(p1)
	printPerson2(&p1)
}
