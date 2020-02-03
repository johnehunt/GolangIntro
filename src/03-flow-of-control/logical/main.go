package main

import (
	"fmt"
)

func main() {
	x := 0
	y := 10

	fmt.Println("Relational Operators")
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x >= y: %t\n", x >= y)
	fmt.Printf("x <= y: %t\n", x <= y)

	flag1 := true
	flag2 := false
	fmt.Println("Logical Operators")
	fmt.Printf("(flag1 && flag2): %t\n", (flag1 && flag2))
	fmt.Printf("(flag1 || flag2): %t\n", (flag1 || flag2))
	fmt.Printf("!(flag1 && flag2): %t\n", !(flag1 && flag2))

	x = 5
	y = 3
	fmt.Println("Artimetic Operators")
	fmt.Printf("x + y: %v\n", x+y)
	fmt.Printf("x - y: %v\n", x-y)
	fmt.Printf("x * y: %v\n", x*y)
	fmt.Printf("x / y: %v\n", x/y)
	fmt.Println("x % y:", (x % y))
	x++
	fmt.Printf("x++: %v\n", x)
	x--
	fmt.Printf("x--: %v\n", x)

	x = 5
	y = 3
	fmt.Println("Binary Operators")
	fmt.Printf("x & y: %v\n", x&y)
	fmt.Printf("x | y: %v\n", x|y)
	fmt.Printf("x ^ y: %v\n", x^y)
	fmt.Printf("x << 2: %v\n", x<<2)
	fmt.Printf("x >> 2: %v\n", x>>2)

}
