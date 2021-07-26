package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")
	var n [10]int /* n is an array of 10 integers - all values initialised to zero */

	/* ALl array values are Zero
	for i := 0; i < 10; i++ {
		fmt.Println(n[i])
	}

	/* initialize elements of array n to 0 */
	for i := 0; i < 10; i++ {
		n[i] = i * 10 /* set element at location i */
	}

	/* output each array element's value */
	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	fmt.Println("---------------")

	// Initialising array at creation time
	numbers := []int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("value of x[%d] = %d\n", i, x)
	}

	fmt.Println("---------------")

	// Multidimensional Array Initialization
	matrix := [3][4]int{
		{0, 1, 2, 3},   /*  initializers for row indexed by 0 */
		{4, 5, 6, 7},   /*  initializers for row indexed by 1 */
		{8, 9, 10, 11}} /*  initializers for row indexed by 2 */

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < (len(matrix[i])); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
		fmt.Println("=")
	}

	fmt.Println("Done")
}
