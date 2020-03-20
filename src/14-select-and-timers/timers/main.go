package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting")

	// Timers represent a single event in the future
	timer := time.NewTimer(2 * time.Second)

	go func() {
		// blocks on the timer’s channel C until it
		// sends a value indicating that the timer fired
		<-timer.C
		fmt.Println("Timer fired")
	}()

	// Try uncommenting  the following
	// stop := timer.Stop()
	// if stop {
	// 	fmt.Println("Timer stopped")
	// }

	time.Sleep(4 * time.Second)

	fmt.Println("Done")
}
