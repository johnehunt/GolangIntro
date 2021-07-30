package main

import (
	"fmt"

	"bank-account/accounts"
)

func main() {
	acc1 := accounts.NewAccount("123", "John", 10.05, "current")
	acc2 := accounts.NewAccount("345", "John", 23.55, "savings")
	acc3 := accounts.NewAccount("567", "Phoebe", 12.45, "investment")

	fmt.Println(acc1)
	fmt.Println(acc2)
	fmt.Println(acc3)

	acc1.Deposit(15.55)
	fmt.Printf("Balance: %6.2f\n", acc1.GetBalance())

	_, err := acc1.Withdraw(120.33)
	if err != nil {
		fmt.Printf("Opps something went wrong: %v\n", err)
	}

	fmt.Printf("Balance now: %6.2f\n", acc1.GetBalance())

}
