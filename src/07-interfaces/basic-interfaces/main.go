package main

import (
	"fmt"
	"math"
)

// Shape basic interface for shape like things
type Shape interface {
	Area() float64
}

// Rectangle (shapes)
type Rectangle struct {
	width, height float64
}

// Circle (shapes)
type Circle struct {
	radius float64
}

// Area To implement an interface in go,
// implement all the methods in the interface
// for a specific type
// Note implementing method uses a value receiver
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Area implements the Shape interface for Cirlce
// Note implementing method uses a value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Can call all methods in the interface Shape
// on the parameter s
func process(s Shape) {
	fmt.Println("Shape: ", s)
	fmt.Println("s.Area()", s.Area())
}

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	// Rectangle implements the interface Shape
	// as there is a method defined for the
	// struct with the name area - note the
	// compiler works this out - its not explicit
	fmt.Println("Process Rectangle")
	process(r)

	fmt.Println("Process Circle")
	process(c)
}
