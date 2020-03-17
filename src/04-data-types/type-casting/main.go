package main

import "fmt"

func main() {
	var total int = 120
	var count int = 5
	var average float32

	average = float32(total) / float32(count)
	fmt.Printf("The average was: %0.2f\n", average)
}
