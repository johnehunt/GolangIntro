package main

import "fmt"

/*
 Sum can take any number of arguments
*/
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Starting")

	fmt.Println("sum():", Sum())
	fmt.Println("sum(1):", Sum(1))
	fmt.Println("sum(1, 2):", Sum(1, 2))
	fmt.Println("sum(1, 2, 3):", Sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println("sum(nums...):", Sum(nums...))

	fmt.Println("Done")
}
