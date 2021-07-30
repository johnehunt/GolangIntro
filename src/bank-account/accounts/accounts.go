package accounts

import (
	"fmt"
	"math"
)

// Private to package variable representing NaN
var nan = math.NaN()

// Public function that provdies 'constant' NaN vlaue
func NaN() float64 {
	return nan
}

// InvalidArgumentError custom exception type
// Implements the Error interface with the Errors method
type InvalidArgumentError struct {
	Argument float64
	Problem  string
}

// Required method (Error interface)
func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("Amount %6.2f - %s", e.Argument, e.Problem)
}

/*
Account represents a type of bank account
*/
type Account struct {
	number      string
	holder      string
	balance     float64
	accountType string
}

/*
NewAccount constructor functions for Account
*/
func NewAccount(number string,
	holder string,
	balance float64,
	accountType string) *Account {

	var acc = Account{number, holder, balance, accountType}
	return &acc

}

// Public Account Updater methods

/*
Deposit add an amount to the Accounts balance
*/
func (acc *Account) Deposit(amount float64) {
	acc.balance = acc.balance + amount
}

/*
Withdraw return an amount from the account.
Returns the new balance or an error
*/
func (acc *Account) Withdraw(amount float64) (float64, error) {
	tempBalance := acc.balance - amount
	if tempBalance < 0 {
		return NaN(), &InvalidArgumentError{amount, "Would take Account balance below Zero"}
	} else {
		acc.balance = tempBalance
		return tempBalance, nil
	}
}

// Public Account Getter methods

func (acc Account) GetNumber() string {
	return acc.number
}

func (acc Account) GetHolder() string {
	return acc.holder
}

/*
GetBalance return the value fo the balance for the account
*/
func (acc Account) GetBalance() float64 {
	return acc.balance
}

func (acc Account) GetType() string {
	return acc.accountType
}

// Implementing the Stringer interface

/*
String convert the Account into a string format
*/
func (acc Account) String() string {
	return fmt.Sprintf("Account(Number: %s, Holder: %s, Balance: %6.2f, Type: %s)",
		acc.number, acc.holder, acc.balance, acc.accountType)
}
