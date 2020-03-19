package main

import (
	"fmt"
	"os"
)

func PanicIfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Logging error: %s\n", err)
	panic(err)
}

func main() {
	filename := "invalid.txt"
	file, err := os.Open(filename)
	PanicIfError(err)
	fmt.Println("file opened", file.Name())
}
