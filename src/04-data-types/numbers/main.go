package main

import "fmt"

func main() {
	var i int
	var c byte
	var salary float32
	flag := true

	i = 5
	fmt.Printf("i is of type %T with value: %d\n", i, i)

	c = 12
	fmt.Printf("c is of type %T with value: %d\n", c, c)

	salary = 25000.00
	fmt.Printf("salary is of type %T with value: %0.2f\n", salary, salary)

	fmt.Printf("flag is of type %T with value: %t\n", flag, flag)

}
