package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// Degree struct
type Degree struct {
	title string
}

// Student struct
type Student struct {
	university string
	Person
	Degree
}

func main() {
	var s = Student{"UWE", Person{"John", 21}, Degree{"Computing"}}
	fmt.Println("s:", s)
}
