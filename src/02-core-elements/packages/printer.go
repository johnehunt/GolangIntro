package main

import (
	"02-core-elements/packages/util"
	"fmt"
)

func main() {
	fmt.Printf("Max value is %d\n", util.MAX)

	temp := util.Freezing
	fmt.Printf("Freezing is %f\n", temp)

	var d util.Distance = 5
	fmt.Printf("The distance is %d\n", d)

}
