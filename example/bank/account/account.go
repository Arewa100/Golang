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
	if validateAmount(amount) && accountNumber == account.accountNumber {
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
