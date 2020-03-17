package main

import "fmt"

func main() {

	var sampleText = "John was here"
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q \n", sampleText)

	var total = 25
	fmt.Printf("Total calculated was = %d\n", total)

	var flag = true
	fmt.Printf("flag: %t\n", flag)

	var rate = 3.45
	fmt.Printf("%f\n", rate)
	fmt.Printf("%0.2f\n", rate)

	// In hex
	fmt.Printf("total %d, in hex: %x\n", total, total)

	// Layout formatting
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("Type of %T", rate)
}
