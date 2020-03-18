package main

import "fmt"

// Person struct has fields Name and Age
// Public structs have an initial capital letter
// and a comment
type Person struct {
	Name string
	Age  int
}

func main() {

	// Declares p1 to be of type Person and
	// initialises the Person struct allocating memory as required
	var p1 Person
	p1.Name = "John"
	p1.Age = 21

	// Can access and update fields using dot notation
	var p2 Person
	p2.Name = "Denise"
	p2.Age = 25

	fmt.Println(p1)
	fmt.Printf("Person p1's Name is: %s\n", p1.Name)
	fmt.Printf("Person p1's Age is: %d\n", p1.Age)

	fmt.Printf("Person p2 %v\n", p2)

	fmt.Println("----------------")

	p3 := Person{Name: "John", Age: 36}
	fmt.Printf("Person p3 %v\n", p3)

	p4 := Person{"Phoebe", 21}
	fmt.Printf("Person p4 %v\n", p4)

	fmt.Println("----------------")

	fmt.Println("Compare structs")
	// Can compare structs with comparison operators
	// such as == and !=

	px1 := Person{"John", 55}
	px2 := Person{"Denise", 53}
	px3 := Person{"John", 55}
	fmt.Println("px1 == px2:", px1 == px2) // false
	fmt.Println("px1 != px2:", px1 != px2) // true
	fmt.Println("px1 == px3:", px1 == px3) // true

	fmt.Println("----------------")

	// Copy struct in p1 to px
	px := p1
	fmt.Printf("Person px's Name is: %s\n", px.Name)
	px.Name = "Adam"
	fmt.Printf("Person px's Name is: %s\n", px.Name)
	fmt.Printf("Person p1's Name is: %s\n", p1.Name)
	fmt.Println(p1)
	fmt.Println(px)

	fmt.Println("----------------")

	// Make a new reference to p1 - An & prefix yields a pointer to the struct
	pref := &p1
	fmt.Printf("Person pref's Name is: %s\n", pref.Name)
	pref.Name = "Paul"
	fmt.Printf("Person pref's Name is: %s\n", pref.Name)
	fmt.Printf("Person pref's Name is: %s\n", pref.Name)

}
