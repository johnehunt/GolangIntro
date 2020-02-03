package main

import (
	"fmt"
)

func main() {

	var x = 1       // Var containing vaklue 1
	var p = &x      // var containing the address of the variable x
	fmt.Println(*p) // Print result of dereferencing p

	*p = 2         // Equal to x = 2
	fmt.Println(x) // will now print 2

	*p++           // Increments what P points to not p
	fmt.Println(x) // will now print 3

	q := new(int) // Creates a pointer to an unnamed int variable
	fmt.Println(*q)
	*q = 2
	fmt.Println(*q)

}
