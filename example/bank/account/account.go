package account

import "time"

type Account struct {
	name           string
	accountNumber  string
	address        string
	dateOfCreation time.Time
	password       string
	balance        float32
}

func CreatAccount(name string, accountNumber string, address string) *Account {
	return &Account{
		name:           name,
		accountNumber:  accountNumber,
		address:        address,
		dateOfCreation: time.Now(),
	}
}
