package main

import "fmt"

// Person struct has fields name and age
type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person /* Declares p1 of type Person */
	p1.name = "John"
	p1.age = 21

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
