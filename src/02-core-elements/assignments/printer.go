package main

import (
	"fmt"
	"strconv"
)

// MAX max value (not type int is optional)
const MAX int = 10

func main() {
	fmt.Println("Starting")
	// Declare a variable but without a value
	// Will get the 'zeroeth' value for the type string
	// which is the empty string ""
	var content string

	// Declare a variable of type int (inferred) with the value 5
	var result = 5
	for i := 1; i < MAX; i++ {
		content = content + "string" + strconv.Itoa(i) + ", "
		result = result * i
	}
	fmt.Println(content)
	fmt.Printf("Total calculated was = %d\n", result)

	fmt.Println("Assignment Operators")

	// Short hand form of declaration and initial vlaue
	x := 5

	// Various Assignment operators
	fmt.Printf("x =5: %v\n", x)
	x += 5
	fmt.Printf("x +=5: %v\n", x)
	x -= 5
	fmt.Printf("x -=5: %v\n", x)
	x *= 5
	fmt.Printf("x *=5: %v\n", x)
	x /= 5
	fmt.Printf("x /=5: %v\n", x)
	x %= 5
	fmt.Println("x %=5:", x)
	x++
	fmt.Println("x++:", x)

}
