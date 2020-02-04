package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "invalid.txt"

	file, err := os.Open(filename)

	//er will be nil if the file exists else it returns an error object
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file opened", file.Name())
	}
}
