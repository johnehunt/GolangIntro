package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// each read we pick a key to access, Lock() the mutex to
// ensure exclusive access to the state, read the value at the
// chosen key, Unlock() the mutex, and increment the readOps count.
func reader(mutex *sync.Mutex, state map[int]int, readOps *uint64) {
	total := 0
	for {
		key := rand.Intn(5)
		mutex.Lock()
		total += state[key]
		mutex.Unlock()
		atomic.AddUint64(readOps, 1)

		time.Sleep(time.Millisecond)
	}
}

// Start 10 goroutines to simulate writes, using
// the same pattern we did for reads.
func writer(mutex *sync.Mutex, state map[int]int, writeOps *uint64) {
	for {
		key := rand.Intn(5)
		val := rand.Intn(100)
		mutex.Lock()
		state[key] = val
		mutex.Unlock()
		atomic.AddUint64(writeOps, 1)
		// Wait a bit between reads.
		time.Sleep(time.Millisecond)
	}
}

func main() {

	// State is a map
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	// keep track of how many read and write operations
	var readOps uint64
	var writeOps uint64

	// start 100 goroutines to execute repeated reads
	for r := 0; r < 100; r++ {
		go reader(mutex, state, &readOps)
	}

	for w := 0; w < 10; w++ {
		go writer(mutex, state, &writeOps)
	}

	// Let the 10 goroutines work on the state and mutex for a second
	time.Sleep(time.Second)

	// Take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Grab the lock to show the state
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
