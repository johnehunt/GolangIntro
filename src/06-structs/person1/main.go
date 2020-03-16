package main

import "fmt"

// Person struct has fields name and age
// Public structs have an initial capital letter
// and a comment
type Person struct {
	name string
	age  int
}

func main() {

	// Declares p1 to be of type Person and
	// initialises the Person struct allocating memory as required
	var p1 Person
	p1.name = "John"
	p1.age = 21

	// Can access and update fields using dot notation
	var p2 Person
	p2.name = "Denise"
	p2.age = 25

	fmt.Println(p1)
	fmt.Printf("Person p1's name is: %s\n", p1.name)
	fmt.Printf("Person p1's age is: %d\n", p1.age)

	fmt.Printf("Person p2 %v\n", p2)

	fmt.Println("----------------")

	var p3 = Person{name: "Phoebe", age: 21}
	fmt.Printf("Person p3 %v\n", p3)

	fmt.Println("----------------")

	fmt.Println("Compare structs")
	// Can compare structs with comparison operators
	// such as == and !=
	fmt.Println("p1 == p2:", p1 == p2) // false
	fmt.Println("p1 != p2:", p1 != p2) // false

	fmt.Println("----------------")

	// Copy struct in p1 to px
	var px = p1
	fmt.Printf("Person px's name is: %s\n", px.name)
	px.name = "Adam"
	fmt.Printf("Person px's name is: %s\n", px.name)
	fmt.Printf("Person p1's name is: %s\n", p1.name)

	fmt.Println("----------------")

	// Make a new reference to p1 - An & prefix yields a pointer to the struct
	var pref = &p1
	fmt.Printf("Person pref's name is: %s\n", pref.name)
	pref.name = "Paul"
	fmt.Printf("Person pref's name is: %s\n", pref.name)
	fmt.Printf("Person pref's name is: %s\n", pref.name)

}
