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

	for i := 0; i < 5; i++ {
		fmt.Printf("value of i: %d\n", i)
		if i >= 3 {
			/* terminate this iteration fo the loop */
			continue
		}
		fmt.Printf("value of i*i: %d\n", i*i)
	}

	fmt.Println("-------------------- for without options")

	j := 0
	for {
		if j > 5 {
			break
		}
		fmt.Printf("j is = %d\n", j)
		j++
	}

	// Using the gogo statement to jump to a label

	fmt.Println("------------------ goto")

	i := 1
LOOP:
	for i < 5 {
		for j := 1; j < 5; j++ {
			if j == 3 {
				i++
				goto LOOP
			}
			fmt.Printf("%v * %v is %v\n", i, j, i*j)
		}
		i++
	}

	fmt.Println("Done")
}
