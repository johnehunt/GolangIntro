package accounts

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	accountNumber := "123"
	accountHolder := "John"
	accountBalance := 12.55
	accountType := "current"
	acc := NewAccount(accountNumber, accountHolder, accountBalance, accountType)
	if acc.GetNumber() != accountNumber {
		t.Error("Expected ", accountNumber, ", got ", acc.GetNumber())
	}
	if acc.GetHolder() != accountHolder {
		t.Error("Expected ", accountHolder, "got ", acc.GetNumber())
	}
	if acc.GetBalance() != accountBalance {
		t.Error("Expected ", accountBalance, "got ", acc.GetBalance())
	}
	if acc.GetType() != accountType {
		t.Error("Expected ", accountHolder, "got ", acc.GetNumber())
	}
}

func TestDeposit10InitialBalance0(t *testing.T) {
	acc := NewAccount("123", "John", 0.0, "current")
	acc.Deposit(10.0)
	balance := acc.GetBalance()
	if balance != 10.0 {
		t.Error("Expected balance to be 10.0, but was ", balance)
	}
}

func TestDeposit10InitialBalance10(t *testing.T) {
	acc := NewAccount("123", "John", 10.0, "current")
	acc.Deposit(10.0)
	balance := acc.GetBalance()
	if balance != 20.0 {
		t.Error("Expected balance to be 20.0, but was ", balance)
	}
}

func TestWithdraw10From10(t *testing.T) {
	acc := NewAccount("123", "John", 10.0, "current")
	acc.Withdraw(10.0)
	balance := acc.GetBalance()
	if balance != 0.0 {
		t.Error("Expected balance to be 0.0, but was ", balance)
	}
}

func TestWithdrawGreaterThanBalance(t *testing.T) {
	acc := NewAccount("123", "John", 10.00, "current")
	_, err := acc.Withdraw(20.00)
	if err == nil {
		t.Error("Expected an Error, but error was nil ")
	}
}
