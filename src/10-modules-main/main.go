package main

import (
	"10-modules/util"
	"fmt"
)

func main() {
	fmt.Printf("Max is %d\n", util.MAX)
	fmt.Printf("Freezing is %0.2f\n", util.Freezing)
	const boiling util.Celsius = 100.0
	fmt.Printf("Boiling is %0.2f\n", boiling)
	var v interface{} = "John"
	util.Printer(v)
	var m = util.Model{"123", "Adam"}
	fmt.Println(m)
}
