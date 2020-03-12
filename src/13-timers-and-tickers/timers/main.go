package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting")

	// Timers represent a single event in the future
	timer1 := time.NewTimer(2 * time.Second)

	// blocks on the timer’s channel C until it
	// sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Example fo cancelling a timmer before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// Try commenting out the following
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(4 * time.Second)

	fmt.Println("Done")
}
