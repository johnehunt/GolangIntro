package main

import "fmt"

const MAX_GOROUTINES = 5

func main() {
	fmt.Println("Starting")
	tasks := make(chan string)
	for i := 0; i < MAX_GOROUTINES; i++ {
		go worker(tasks)
	}
	tasks <- "Do Work1"
	tasks <- "Do Work2"
	tasks <- "Do Work3"
	fmt.Println("Done")
}

func worker(tasks chan string) {
	fmt.Println("worker > received", <-tasks)
}
