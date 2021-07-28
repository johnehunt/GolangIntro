package main

import (
	"fmt"
	"sync"
	"time"
)

// Set up local variables
var (
	rwMutex sync.RWMutex // Multiple readers, single writer lock
	balance int
)

func Deposit(amount int) {
	rwMutex.Lock()         // Get a Write lock
	defer rwMutex.Unlock() // Release Write Lock when func terminates
	balance = balance + amount
}

func Withdraw(amount int) {
	rwMutex.Lock()         // Get a Write lock
	defer rwMutex.Unlock() // Release Write Lock when func terminates
	balance = balance - amount
}

func Balance() int {
	rwMutex.RLock()         // Get read lock
	defer rwMutex.RUnlock() // release read lock when function terminates
	return balance
}

func main() {
	fmt.Println("Starting")

	fmt.Printf("Initial balance is %d\n", Balance())

	go func() {
		Deposit(200)
		fmt.Printf("Balance after 200 deposit is %d\n", Balance())
	}()

	go func() {
		Withdraw(150)
		fmt.Printf("Balance after 150 withdrawal is %d\n", Balance())
	}()

	go func() {
		Deposit(100)
		fmt.Printf("Balance after 100 deposit is %d\n", Balance())
	}()

	time.Sleep(2 * time.Second)

	fmt.Printf("Final balance is %d\n", balance)

	fmt.Println("Done")
}
