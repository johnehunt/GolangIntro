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

// Set up bank account and channels for comms
type Account struct {
	balance         float64
	depositChannel  chan float64
	withdrawChannel chan float64
	balanceChannel  chan float64
}

func NewAccount(initialBalance float64) *Account {
	acc := Account{
		balance:         initialBalance,
		depositChannel:  make(chan float64),
		withdrawChannel: make(chan float64),
		balanceChannel:  make(chan float64),
	}
	return &acc
}

func (acc *Account) Deposit(amount float64) {
	acc.depositChannel <- amount
}

func (acc *Account) Withdraw(amount float64) {
	acc.withdrawChannel <- amount
}

func (acc *Account) Balance() float64 {
	return <-acc.balanceChannel
}

func (acc *Account) monitor() {
	for {
		select {
		case amount := <-acc.depositChannel:
			acc.balance = acc.balance + amount
		case amount := <-acc.withdrawChannel:
			acc.balance = acc.balance - amount
		case acc.balanceChannel <- acc.balance:
		}
	}
}

func main() {

	fmt.Println("Starting")
	// Create the account
	acc := NewAccount(0.0)

	// Start the monitor running in the background
	go acc.monitor()

	fmt.Println("Starting the go routines")
	go func() {
		acc.Deposit(200)
		fmt.Printf("Balance after 200 deposit is %4.2f\n", acc.Balance())
	}()

	go func() {
		acc.Withdraw(150)
		fmt.Printf("Balance after 150 withdrawal is %4.2f\n", acc.Balance())
	}()

	go func() {
		acc.Deposit(100)
		fmt.Printf("Balance after 100 deposit is %4.2f\n", acc.Balance())
	}()

	time.Sleep(2 * time.Second)
	fmt.Printf("Final balance is %4.2f\n", acc.Balance())

	fmt.Println("Done")
}

/*
Sample Output:
--------------

Starting
Starting the go routines
Balance after 200 deposit is 300.00
Balance after 150 withdrawal is 150.00
Balance after 100 deposit is 100.00
Final balance is 150.00
Done

*/
