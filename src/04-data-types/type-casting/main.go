// Note Go doesn’t support implicit type conversion
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var total int = 120
	var count int = 5
	var average float32

	average = float32(total) / float32(count)
	fmt.Printf("The average is: %0.2f\n", average)

	i := int(average)
	fmt.Printf("The integer average is: %d\n", i)

	// Converting a String to an int
	var s string = "42"
	v, e := strconv.Atoi(s)

	if e != nil {
		fmt.Printf("Something when wrong %v\n", e)
	} else {
		fmt.Printf("The string as an int is %d\n", v)
	}

	// Converting an int to a String
	age := 42
	str := strconv.Itoa(age) // convert int to string
	fmt.Printf("The int as a string is: %s\n", str)

	// Converting a string to a float
	f := "3.14159265"
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s) // 3.14159265
	}

	// Convert a float to a string
	str2 := fmt.Sprintf("%4.2f", 123.456)
	fmt.Println(str2)

}
