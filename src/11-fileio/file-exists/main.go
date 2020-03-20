package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	fmt.Println("File does exist. File information:")
	fmt.Println(fileInfo)
}
