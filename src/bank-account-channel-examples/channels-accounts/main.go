package main

import (
	"fmt"
	"time"
)

/*
All concurrent access with the balance is via channels
all handled via a central monitor which ensures sequential
access once requests are sent via the channel (which can run concurrently)
*/

// Set up channels for comms
var depositChannel = make(chan int)
var withdrawChannel = make(chan int)
var balanceChannel = make(chan int)

func Deposit(amount int) {
	depositChannel <- amount
}

func Withdraw(amount int) {
	withdrawChannel <- amount
}

func Balance() int {
	return <-balanceChannel
}

func monitor() {
	// balance is lexically confined to function
	var balance int
	for {
		select {
		case amount := <-depositChannel:
			balance = balance + amount
		case amount := <-withdrawChannel:
			balance = balance - amount
		case balanceChannel <- balance:
		}
	}
}

func init() {
	go monitor() // start the balance monitor goroutine
}

func main() {
	fmt.Println("Starting")

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
	fmt.Printf("Final balance is %d\n", Balance())

	fmt.Println("Done")
}

/*
Sample Output:
--------------

Starting
Balance after 200 deposit is 200
Balance after 150 withdrawal is 50
Balance after 100 deposit is 150
Final balance is 150
Done

*/
