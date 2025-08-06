package account

import (
	"errors"
	"time"
)

type Account struct {
	name           string
	accountNumber  string
	address        string
	dateOfCreation time.Time
	password       string
	balance        float32
}

func CreateAccount(name string, accountNumber string, address string, password string) *Account {
	return &Account{
		name:           name,
		accountNumber:  accountNumber,
		address:        address,
		dateOfCreation: time.Now(),
		password:       password,
		balance:        0.0,
	}
}
func (account *Account) CheckAccountBalance(accountNumber string, password string) float32 {
	return account.balance
}

func (account *Account) Deposit(accountNumber string, amount float32) error {
	if !validateAmount(amount) {
		return errors.New("invalid amount")
	}
	if accountNumber == account.accountNumber {
		account.balance += amount
		return nil
	} else {
		return errors.New("account number is not valid")
	}
}

func validateAmount(amount float32) bool {
	if amount > 0.0 {
		return true
	} else {
		return false
	}
}

func (account *Account) Withdraw(accountNumber string, amount float32, password string) error {
	if !validateAmount(amount) {
		return errors.New("invalid withdraw amount")
	}
	if accountNumber == account.accountNumber && password == account.password {
		if account.balance > amount {
			account.balance -= amount
			return nil
		} else {
			return errors.New("insufficient balance")
		}
	} else {
		return errors.New("invalid account number or password")
	}
}
