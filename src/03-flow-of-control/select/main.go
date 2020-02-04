package main

import (
	"fmt"
)

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received %v from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %v to c2\n", i2)
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received %v from c3\n", i3)
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("No communication")
	}
}
