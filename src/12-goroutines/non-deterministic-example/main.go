package main

import "fmt"

func main() {
	n := 10
	for i := 0; i < n; i++ {
		go println(i)
	}
	fmt.Println("Done")
	// when is everything done?
}
