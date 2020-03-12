package main

import (
	"fmt"
	"math"
)

// Shape basic interface for chapes
type Shape interface {
	area() float64
}

// Rectangle (shapes)
type Rectangle struct {
	width, height float64
}

// Circle (shapes)
type Circle struct {
	radius float64
}

// To implement an interface in go,
// implement all the methods in the interface
// for a specific type
// Note implementing method uses a value receiver
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// This implements the SHape interface for Cirlce
// Note implementing method uses a value receiver
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Can call all methods in the interface Shape
// on the parameter s
func process(s Shape) {
	fmt.Println("Shape: ", s)
	fmt.Println("s.area()", s.area())
}

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	// Rectangle implement the interface Shape
	// as there is a method defined for the
	// struct with the name area
	fmt.Println("Process Rectangle")
	process(r)

	fmt.Println("Process Circle")
	process(c)
}
