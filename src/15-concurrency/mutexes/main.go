package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Each read pick a key to access, Lock() the mutex to
// ensure exclusive access to the state and read the value at the
// chosen key, Unlock() the mutex
func reader(mutex *sync.RWMutex, state map[int]int) {
	for {
		key := rand.Intn(5)
		mutex.RLock() // Lock for reading
		fmt.Printf("Reading key %d - %d\n", key, state[key])
		mutex.RUnlock() // Unlock for reading (could use defer as safer)
		time.Sleep(time.Millisecond)
	}
}

// Start writer goroutine
func writer(mutex *sync.RWMutex, state map[int]int) {
	for {
		key := rand.Intn(5)
		value := rand.Intn(100)
		mutex.Lock() // Lock for writing
		fmt.Printf("Writing key %d - %d\n", key, value)
		state[key] = value
		mutex.Unlock() // Unlock for writing (could use defer as safer)
		// Wait a bit between writes.
		time.Sleep(time.Millisecond)
	}
}

func main() {
	var state = make(map[int]int)
	var mutex = &sync.RWMutex{}

	// start 100 goroutines to execute repeated reads
	for r := 0; r < 100; r++ {
		go reader(mutex, state)
	}

	// Start 10 writer goroutines
	for w := 0; w < 10; w++ {
		go writer(mutex, state)
	}

	// Let the 10 goroutines work on the state and mutex for a second
	time.Sleep(time.Second)

	// Grab the lock to show the state
	mutex.RLock()
	fmt.Println("state:", state)
	mutex.RUnlock()
}
