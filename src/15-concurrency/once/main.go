package main

import (
	"fmt"
	"sync"
	"time"
)

func setup() {
	fmt.Println("setup only runs once")
}

func main() {
	// Create the once instance
	var once sync.Once

	fmt.Println("Starting")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Calling setup")
			once.Do(setup) // Ensures some behaviour is only run once - ever!
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Done")
}

/*
Sample Output
=============
Starting
Calling setup
Calling setup
Calling setup
setup only runs once
Calling setup
Calling setup
Done
*/
