package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Degree interface
type Degree interface {
	Title() string
}

// Student struct
type Student struct {
	university string
	Person
	Degree
}

// UndergraduateDegree struct implementing the Degree interface
type UndergraduateDegree struct {
	Name string
}

func (d UndergraduateDegree) Title() string {
	return d.Name
}

func main() {
	var s = Student{"UWE", Person{"John", 21}, UndergraduateDegree{"Computing"}}
	fmt.Println("s:", s)
	fmt.Println("s.Title():", s.Title())
}
