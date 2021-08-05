package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for j := range tasks {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("Start")

	// Set up channels
	const TASK_NUMBER = 5
	tasksChannel := make(chan int, TASK_NUMBER)
	resultsChannel := make(chan int, TASK_NUMBER)

	// Start up 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, tasksChannel, resultsChannel)
	}

	// Send 5 tasks (numbered 1 to 5) to the
	// task channel to be processed by the workers
	for j := 1; j <= TASK_NUMBER; j++ {
		tasksChannel <- j
	}

	// Close the task channel to indicate
	// that there is no further work to do
	close(tasksChannel)

	// Collect the results of the work
	for a := 1; a <= TASK_NUMBER; a++ {
		var r = <-resultsChannel
		fmt.Println("result:", r)
	}

	fmt.Println("Done")
}
