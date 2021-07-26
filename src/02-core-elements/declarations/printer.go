package main

import (
	"fmt"
	"strconv"
)

// MAX a max value at top level
const MAX int = 10

func main() {
	// short hand form of declaring var of type string
	// note use of declaration and assignment operator ':='
	content := ""

	// short hand for of declaring result as variable of typ eint
	result := 5
	for i := 1; i < MAX; i++ {
		// Note use of assignment operator '='
		content = content + "string" + strconv.Itoa(i) + ", "
		result = result * i
	}

	fmt.Println(content)
	fmt.Printf("Total calculated was = %d", result)

	// result := 0 // Produces an error as already declared
}
