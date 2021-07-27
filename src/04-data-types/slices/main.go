package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// Create a slice on a backing integer array
	// Slices are a bit like ArrayList or Vecotr in other languages
	// numbers is a slice holding integers of size 3 with capacity 5
	// currently filled with zeros
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	// An empty slice
	var numbers2 []int
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers2), cap(numbers2), numbers2)
	if numbers2 == nil {
		fmt.Println("slice is nil")
	}

	fmt.Println("-------------")
	// Example of using append
	// Note append may return a new slice as data is currently empty
	var data []int
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(data), cap(data), data)
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
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(marks), cap(marks), marks)

	/* print the sub slice starting from index 1(included)
	// to index 4(excluded)*/
	fmt.Println("marks[1:4] ==", marks[1:4])

	/* missing lower bound implies 0 */
	fmt.Println("marks[:3] ==", marks[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("marks[4:] ==", marks[4:])

	fmt.Println("-------------")

	/* subsclice is a view onto same array as original slice */
	fmt.Println(marks)
	tempMarks := marks[1:4]
	tempMarks[1] = 100
	fmt.Println(marks)

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
