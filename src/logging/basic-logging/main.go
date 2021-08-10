package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)

	log.Print("Logging in Go!")
	log.Println("Logging with a newline")
	log.Printf("Logging using printf formatting %s", "hello")
}
