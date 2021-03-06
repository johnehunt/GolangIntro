package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	var numbers2 []int
	if numbers2 == nil {
		fmt.Println("slice is nil")
	}

	fmt.Println("-------------")
	// Append
	var data []int
	data = append(data, 1)
	fmt.Println("data =", data)
	fmt.Printf("data len=%d, cap=%d\n", len(data), cap(data))

	data = append(data, 2, 3, 4)
	fmt.Println("data =", data)
	fmt.Printf("data len=%d, cap=%d\n", len(data), cap(data))

	fmt.Println("-------------")

	// Subslicing
	marks := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	/* print the original */
	fmt.Println("marks ==", marks)

	/* print the sub slice starting from index 1(included)
	// to index 4(excluded)*/
	fmt.Println("marks[1:4] ==", marks[1:4])

	/* missing lower bound implies 0 */
	fmt.Println("marks[:3] ==", marks[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("marks[4:] ==", marks[4:])

	fmt.Println("-------------")

	// Copy

	fmt.Printf("data = %v\n", data)
	/* create a slice numbers1 with double the capacity of earlier slice*/
	data1 := make([]int, len(data), (cap(data))*2)

	/* copy content of numbers to numbers1 */
	copy(data1, data)
	fmt.Println("data1 =", data1)

	fmt.Printf("data len=%d cap=%d slice=%v\n", len(data), cap(data), data)
	fmt.Printf("data1 len=%d cap=%d slice=%v\n", len(data1), cap(data1), data1)

	fmt.Println("-------------")

	// Remove Elements
	scores := []int{10, 11, 12, 13}
	fmt.Printf("scores len=%d cap=%d slice=%v\n", len(scores), cap(scores), scores)
	index := 1 // index of 11
	scores = append(scores[:index], scores[index+1:]...)
	fmt.Printf("scores len=%d cap=%d slice=%v\n", len(scores), cap(scores), scores)

	fmt.Println("Done")
}
