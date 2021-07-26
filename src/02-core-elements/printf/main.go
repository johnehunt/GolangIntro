package main

import "fmt"

func main() {

	var sampleText = "John was here"
	fmt.Printf("plan string: %s \n", sampleText)
	fmt.Printf("quoted string: %+q \n", sampleText)

	total := 25
	fmt.Printf("Total calculated was = %d\n", total)

	flag := true
	fmt.Printf("flag: %t\n", flag)

	const RATE = 3.45
	fmt.Printf("%f\n", RATE)
	fmt.Printf("%0.2f\n", RATE)

	// In hex
	fmt.Printf("total %d, in hex: %x\n", total, total)

	// Layout formatting
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("Type of RATE: %T\n", RATE)
}
