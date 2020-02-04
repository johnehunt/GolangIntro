package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// Using the break statement

	for i := 0; i < 10; i++ {
		fmt.Printf("value of i: %d\n", i)
		if i == 5 {
			/* terminate the loop using break statement */
			break
		}
	}

	fmt.Println("------------------")

	// Using the continue statement

	for i := 0; i < 10; i++ {
		fmt.Printf("value of i: %d\n", i)
		if i >= 5 {
			/* terminate this iteration fo the loop */
			continue
		}
		fmt.Printf("value of i*i: %d\n", i*i)
	}

	// Using the gogo statement to jump to a label

	fmt.Println("------------------ goto")

	i := 0
	LOOP: for i < 10 {
		for j := 0; j < 10; j++ {
			if j == 5 {
				i++
				goto LOOP
			}
			fmt.Printf("%v * %v is %v\n", i, j, i*j)
		}
		i++
	}

	fmt.Println("Done")
}
