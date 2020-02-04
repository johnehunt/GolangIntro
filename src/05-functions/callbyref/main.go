package main

import "fmt"

/* function definition to swap the values */
func swap(x *int, y *int) {
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

	swap(&a, &b)

	fmt.Printf("After Swap\na is %d, b os %d\n", a, b)

	fmt.Println("Done")
}
