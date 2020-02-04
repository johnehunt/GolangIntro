package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Starting")

	fmt.Println("sum(1, 2):", sum(1, 2))
	fmt.Println("sum(1, 2, 3):", sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println("sum(nums...):", sum(nums...))

	fmt.Println("Done")
}
