package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for j := range tasks {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("Start")

	// Set up channels
	const taskNumber = 5
	tasks := make(chan int, taskNumber)
	results := make(chan int, taskNumber)

	// Start up 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, tasks, results)
	}

	// Send 5 tasks (numbered 1 to 5) to the
	// task channel to be processed by the workers
	for j := 1; j <= taskNumber; j++ {
		tasks <- j
	}

	// Close that channel to indicate that’s all the work we have.
	close(tasks)

	// Collect the results of the work
	for a := 1; a <= taskNumber; a++ {
		var r = <-results
		fmt.Println("result:", r)
	}

	fmt.Println("Done")
}
