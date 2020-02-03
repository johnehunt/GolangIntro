package main

import (
	"fmt"
	"strconv"
)

const MAX int = 10

func main() {
	var content string
	var result = 5
	for i := 1; i < MAX; i++ {
		content = content + "string" + strconv.Itoa(i) + ", "
		result = result * i
	}
	fmt.Println(content)
	fmt.Printf("Total calculated was = %d", result)
}
