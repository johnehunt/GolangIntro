package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	// Basic Rate Limiting
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limiter channel will receive a value every 200 milliseconds
	// Acts as the regulator in rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// Blocks on receive from limiter channel before serving each request,
	// Thus limited to handling 1 request very 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("Done")
}
