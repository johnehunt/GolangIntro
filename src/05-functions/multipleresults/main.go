package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Starting")

	s1 := "Hello"
	s2 := "World"

	fmt.Println("s1, s2:", s1, s2)
	s3, s4 := swap(s1, s2)
	fmt.Println("s3, s4:", s3, s4)

	fmt.Println("Done")
}
