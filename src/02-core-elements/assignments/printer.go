package main

import (
	"fmt"
	"strconv"
)

// MAX max value
const MAX int = 10

func main() {
	var content string
	var result = 5
	for i := 1; i < MAX; i++ {
		content = content + "string" + strconv.Itoa(i) + ", "
		result = result * i
	}
	fmt.Println(content)
	fmt.Printf("Total calculated was = %d", result)

	fmt.Println("Assignment Operators")

	var x = 5
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

}
