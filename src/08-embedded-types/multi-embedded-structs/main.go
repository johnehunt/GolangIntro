package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Degree struct
type Degree struct {
	Name string
	Type string
}

// Student struct
type Student struct {
	University string
	Person
	Degree
}

func main() {
	var s = Student{"UWE", Person{"John", 21}, Degree{"Computing", "BSc"}}
	fmt.Println("s:", s)

	fmt.Println("s.Type:", s.Type)
	// fmt.Println("s.Name:", s.Name) // won't compile ambiguous selector s.Name
	fmt.Println("s.Person.Name:", s.Person.Name)
}
