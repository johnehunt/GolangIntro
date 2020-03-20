package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Starting")

	rune1 := 'A'
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s\n",
		rune1, rune1, reflect.TypeOf(rune1))

	rune2 := '\u266C'
	fmt.Printf("Rune 2: %c; Unicode: %U; Type: %s\n",
		rune2, rune2, reflect.TypeOf(rune2))

	fmt.Println("Done")
}
