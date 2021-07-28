package main

import (
	"fmt"
	"time"
)

// Is this code concurrency-safe?
// Simple example of how not to access a shared writable variable

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Withdraw(amount int) {
	balance = balance - amount
}

func main() {
	fmt.Println("Starting")

	fmt.Printf("Initial balance is %d\n", balance)

	go func() {
		Deposit(200)
		fmt.Printf("Balance after 200 deposit is %d\n", balance)
	}()

	go func() {
		Withdraw(150)
		fmt.Printf("Balance after 150 withdrawal is %d\n", balance)
	}()

	go func() {
		Deposit(100)
		fmt.Printf("Balance after 100 deposit is %d\n", balance)
	}()

	time.Sleep(2 * time.Second)

	fmt.Printf("Final balance is %d\n", balance)

	fmt.Println("Done")
}

/*
Sample output:
--------------

Starting
Initial balance is 0
Balance after 100 deposit is 150
Balance after 150 withdrawal is 50
Balance after 200 deposit is 200
Final balance is 150
Done

*/
