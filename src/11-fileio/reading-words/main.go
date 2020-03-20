package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file and create scanner on top of it
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines.
	// ScanWords picks out individual elements
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanWords)

	// Scan for next token.
	success := scanner.Scan()
	for success {
		fmt.Println("word:", scanner.Text())
		success = scanner.Scan()
	}
	if success == false {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed")
		} else {
			log.Fatal(err)
		}
	}
}
