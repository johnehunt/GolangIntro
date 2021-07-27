package main

import "fmt"

// Swap swaps two integers using call by reference
func Swap(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}

func main() {
	fmt.Println("Starting")

	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before Swap\na is %d, b is %d\n", a, b)

	Swap(&a, &b)

	fmt.Printf("After Swap\na is %d, b os %d\n", a, b)

	fmt.Println("Done")
}
