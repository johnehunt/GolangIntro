package main

import (
	"fmt"
	"os"
)

func main() {
	// Process each of the command line arguments in turn
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}

// Can run from command line as:
// go run hello2.go John 1 True 42.5
//
// Output will be:
// John
// 1
// True
// 42.5
