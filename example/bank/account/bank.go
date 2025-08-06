package account

import (
	"fmt"
	"time"
)

type Bank struct {
	name           string
	location       string
	dateOfCreation time.Time
	listOfAccount  map[string]Account
}

func CreateBank(name string, location string) *Bank {
	return &Bank{name, location, time.Now(), make(map[string]Account)}
}

func (bank *Bank) CreateAnAccount(name string, accountNumber string, address string, password string) *Account {
	newAccount := CreateAccount(name, accountNumber, address, password)
	accountNumberAsKey := newAccount.accountNumber
	bank.listOfAccount[accountNumberAsKey] = *newAccount
	return newAccount
}

func (bank *Bank) countNumberOfAccounts() float64 {
	return float64(len(bank.listOfAccount))
}

func (bank *Bank) Deposit(accountNumber string, amount float32) error {
	accountToDeposit := bank.getAccount(accountNumber)
	err := accountToDeposit.Deposit(accountNumber, amount)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (bank *Bank) getAccount(accountNumber string) *Account {
	theAccount := bank.listOfAccount[accountNumber]
	return &theAccount
}

func (bank *Bank) CheckBalance(accountNumber string, password string) float32 {
	theAccount := bank.getAccount(accountNumber)
	fmt.Println(theAccount.CheckAccountBalance(accountNumber, "password"))
	fmt.Println(bank.listOfAccount)
	fmt.Println(theAccount)
	balance := theAccount.CheckAccountBalance(accountNumber, password)
	fmt.Println(balance)
	return balance
}
