package main

import (
	"fmt"
	"strconv"
)

// MAX a max value
const MAX int = 10

func main() {
	content := ""
	result := 5
	for i := 1; i < MAX; i++ {
		content = content + "string" + strconv.Itoa(i) + ", "
		result = result * i
	}
	fmt.Println(content)
	fmt.Printf("Total calculated was = %d", result)

	// result := 0 // Produces an error as already declared
}
