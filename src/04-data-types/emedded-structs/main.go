package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Centre Point
	Radius int
}

type Wheel struct {
	circle Circle
	Spokes int
}

// Structs with anonymous fields

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func main() {
	fmt.Println("Starting")

	var wheel1 Wheel
	wheel1.circle.Centre.X = 10
	wheel1.circle.Centre.Y = 10
	wheel1.circle.Radius = 5
	wheel1.Spokes = 10
	fmt.Println(wheel1)

	var wheel2 Wheel2
	wheel2.X = 10
	wheel2.Y = 10
	wheel2.Radius = 5
	wheel2.Spokes = 10
	fmt.Println(wheel2)

	fmt.Println("Done")
}
