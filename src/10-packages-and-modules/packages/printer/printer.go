package main

import (
	"10-packages-and-modules/packages/util"
	"fmt"
)

// Can give import an alias
// import u "10-packages-and-modules/packages/util"

func main() {
	fmt.Printf("Max value is %d\n", util.MAX)

	temp := util.Freezing
	fmt.Printf("Freezing is %f\n", temp)

	var d util.Distance = 5
	fmt.Printf("The distance is %d\n", d)

}
