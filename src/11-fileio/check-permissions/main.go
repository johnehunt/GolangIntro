package main

import (
	"log"
	"os"
)

func main() {
	// Test write permissions
	file, err := os.OpenFile("input.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	// Test read permissions
	file, err = os.OpenFile("input.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}
