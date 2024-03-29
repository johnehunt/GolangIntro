package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	// tickers are for when you want to do something
	// repeatedly at regular intervals
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			t := <-ticker.C
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(3 * time.Second)

	// Tickers can be stopped.
	// Once a ticker is stopped it won’t
	// receive any more values on its channel.
	ticker.Stop()

	fmt.Println("Ticker stopped")

	fmt.Println("Done")
}
