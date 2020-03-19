package main

import (
	"fmt"
	"sort"
)

type LengthSortedStringArray []string

func (s LengthSortedStringArray) Len() int {
	return len(s)
}
func (s LengthSortedStringArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s LengthSortedStringArray) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fmt.Println("Start")
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(LengthSortedStringArray(fruits))
	fmt.Println(fruits)
	fmt.Println("Done")
}
