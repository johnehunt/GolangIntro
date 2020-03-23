package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	f1 := func() {
		fmt.Println("Only once")
	}

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Calling f1")
			once.Do(f1)
		}()
	}
	time.Sleep(time.Second)
}
