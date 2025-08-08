package account

import (
	"errors"
	"time"
)

type Bank struct {
	name           string
	location       string
	dateOfCreation time.Time
	listOfAccount  map[string]*Account
}

func CreateBank(name string, location string) *Bank {
	return &Bank{name, location, time.Now(), make(map[string]*Account)}
}

func (bank *Bank) CreateAnAccount(name string, accountNumber string, address string, password string) *Account {
	newAccount := CreateAccount(name, accountNumber, address, password)
	accountNumberAsKey := newAccount.accountNumber
	bank.listOfAccount[accountNumberAsKey] = newAccount
	return newAccount
}

func (bank *Bank) countNumberOfAccounts() float64 {
	return float64(len(bank.listOfAccount))
}

func (bank *Bank) Deposit(accountNumber string, amount float32) error {
	accountToDeposit := bank.getAccount(accountNumber)
	if bank.checkIfAccountExist(accountNumber) == nil {
		return errors.New("account not found")
	}
	err := accountToDeposit.Deposit(accountNumber, amount)
	if err != nil {
		return err
	}
	return nil
}

func (bank *Bank) getAccount(accountNumber string) *Account {
	theAccount := bank.listOfAccount[accountNumber]
	return theAccount
}

func (bank *Bank) checkIfAccountExist(accountNumber string) error {
	theAccount := bank.listOfAccount[accountNumber]
	if theAccount == nil {
		return errors.New("account not found")
	}
	return nil

	//i am here 
}

func (bank *Bank) CheckBalance(accountNumber string, password string) float32 {
	theAccount := bank.getAccount(accountNumber)
	balance := theAccount.CheckAccountBalance(accountNumber, password)
	return balance
}

func (bank *Bank) WithDraw(accountNumber string, amount float32, password string) error {
	theAccount := bank.getAccount(accountNumber)
	err := theAccount.Withdraw(accountNumber, amount, password)
	if err != nil {
		return err
	}
	return nil
}

func (bank *Bank) Transfer(senderAccount string, recipientAccount string, amount float32, password string) error {
	sender := bank.getAccount(senderAccount)
	receiver := bank.getAccount(recipientAccount)
	err := sender.Withdraw(senderAccount, amount, password)
	if err != nil {
		return err
	}
	receiverError := receiver.Deposit(recipientAccount, amount)
	if receiverError != nil {
		return receiverError
	}
	return nil
}
