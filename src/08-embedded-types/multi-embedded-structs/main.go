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

// Student struct with embedded Person and Degree structs
type Student struct {
	University string
	Person
	Degree
}

func NewStudent(university string, name string, age int, degreeTitle string, degreeType string) *Student {
	student := Student{university, Person{name, age}, Degree{degreeTitle, degreeType}}
	return &student
}

func main() {
	var s = Student{"UWE", Person{"John", 21}, Degree{"Computing", "BSc"}}
	fmt.Printf("s: %v", s)

	fmt.Println("s.Type:", s.Type)
	// fmt.Println("s.Name:", s.Name) // won't compile ambiguous selector s.Name
	fmt.Println("s.Person.Name:", s.Person.Name)

	// Can use a constructor function to splify instantiation
	var s2 = NewStudent("UWE", "Bill", 21, "Cybersecurity", "BSc")
	fmt.Println(s2)

}
