package main

import (
	"fmt"

	"10-modules/calc"
	"10-modules/util"
)

func main() {

	fmt.Printf("Max is %d\n", util.MAX)
	fmt.Printf("Freezing is %0.2f\n", util.Freezing)
	const boiling util.Celsius = 100.0
	fmt.Printf("Boiling is %0.2f\n", boiling)
	var v interface{} = "John"
	util.Printer(v)
	var m = util.Model{ID: "123", Name: "Adam"}
	fmt.Println(m)

	var result, err = calc.Calculator("+", 2, 3)
	if err != nil {
		fmt.Println("Something went wrong", err)
	} else {
		fmt.Println("Result:", result)
	}
}
