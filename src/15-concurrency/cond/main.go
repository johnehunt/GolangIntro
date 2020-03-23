package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, cond *sync.Cond) {
	fmt.Println(id, " - Starting")
	fmt.Println(id, " - Waiting for Cond")
	cond.L.Lock()
	cond.Wait()
	fmt.Println(id, " - Notified by Cond")
	cond.L.Unlock()
	fmt.Println(id, " - Done")
}

func notifier(cond *sync.Cond) {
	fmt.Println("Notifier - Sleeping")
	time.Sleep(time.Second)
	cond.L.Lock()
	fmt.Println("Notifier - Broadcasting notification")
	cond.Broadcast()
	cond.L.Unlock()
	fmt.Println("Notifier - Done")
}

func main() {
	fmt.Println("Starting Worker Goroutines")
	cond := sync.NewCond(&sync.RWMutex{})
	for i := 0; i < 4; i++ {
		go worker(i, cond)
	}
	fmt.Println("Starting Notifier")
	go notifier(cond)

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
