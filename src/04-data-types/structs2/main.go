package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Circle struct {
	center Point
	radius int
}

type Wheel struct {
	circle Circle
	spokes int
}

// Strutcs with anonymous fields
type Circle2 struct {
	Point
	radius int
}

type Wheel2 struct {
	Circle2
	spokes int
}

func main() {
	fmt.Println("Starting")

	var wheel1 Wheel
	wheel1.circle.center.x = 10
	wheel1.circle.center.y = 10
	wheel1.circle.radius = 5
	wheel1.spokes = 10
	fmt.Println(wheel1)

	var wheel2 Wheel2
	wheel2.x = 10
	wheel2.y = 10
	wheel2.radius = 5
	wheel2.spokes = 10
	fmt.Println(wheel2)

	fmt.Println("Done")
}
