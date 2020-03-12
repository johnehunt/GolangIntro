package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// Degree struct
type Degree interface {
	title() string
}

// UndergraduateDegree struct implementing the Degree interface
type UndergraduateDegree struct {
	name string
}

func (d UndergraduateDegree) title() string {
	return d.name
}

// Student struct
type Student struct {
	university string
	Person
	Degree
}

func main() {
	var s = Student{"UWE", Person{"John", 21}, UndergraduateDegree{"Computing"}}
	fmt.Println("s:", s)
}
