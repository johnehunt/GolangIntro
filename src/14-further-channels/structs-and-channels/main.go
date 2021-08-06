package main

import (
	"fmt"
	"time"
)

type Message struct {
	ID     int
	Amount float64
}

type Account struct {
	holder          string
	balance         float64
	depositChannel  chan Message
	withdrawChannel chan Message
	balanceChannel  chan float64
	stop            chan struct{}
}

func (a *Account) Deposit(msg Message) {
	a.depositChannel <- msg
}

func (a *Account) Withdraw(msg Message) {
	a.withdrawChannel <- msg
}

func (a *Account) Balance() float64 {
	return <-a.balanceChannel
}

func (a *Account) Stop() {
	close(a.stop)
}

func (a *Account) monitor() {
	for {
		select {
		case msg := <-a.depositChannel:
			fmt.Println("Received msg:", msg)
			a.balance = a.balance + msg.Amount
		case msg := <-a.withdrawChannel:
			fmt.Println("Received msg:", msg)
			a.balance = a.balance - msg.Amount
		case a.balanceChannel <- a.balance:
			fmt.Println("Received request for balance")
		case <-a.stop:
			fmt.Println("Received stop request")
			return
		}
	}
}

func newAccount(holder string, balance float64) *Account {
	acc := Account{holder: holder,
		balance:         balance,
		depositChannel:  make(chan Message),
		withdrawChannel: make(chan Message),
		balanceChannel:  make(chan float64),
		stop:            make(chan struct{}),
	}
	go acc.monitor() // start the monitor loop
	return &acc
}

var account = newAccount("Bill Smith", 0.0)

func main() {
	fmt.Println("Starting")

	go func() {
		account.Deposit(Message{1, 200.0})
	}()

	go func() {
		account.Withdraw(Message{2, 150.0})
	}()

	go func() {
		account.Deposit(Message{3, 10.0})
	}()

	time.Sleep(time.Second)
	fmt.Printf("Balance is %4.2f\n", account.Balance())

	account.Stop()

	fmt.Println("Done")
}
