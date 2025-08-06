package account

import (
	"fmt"
	"testing"
)

func bank() *Bank {
	koloBank := CreateBank("Kolo", "Chemist bus-stop")
	return koloBank
}
func TestToCreateAccount(t *testing.T) {
	koloBank := bank()
	if koloBank == nil {
		t.Error("error creating kolo bank")
	}
	if koloBank != nil && koloBank.countNumberOfAccounts() != 0 {
		t.Error("error in the kolo bank map of accounts")
	}
	newAccount := koloBank.CreateAnAccount("Ola", "123456", "bajulaye number 1", "password")
	if koloBank != nil && koloBank.countNumberOfAccounts() != 1 {
		t.Error("Expected number of accounts to be 1")
	}

	fmt.Println(newAccount)
}

func TestThatBankCanDeposit(t *testing.T) {
	koloBank := bank()
	theCreatedAccount := koloBank.CreateAnAccount("Ola", "123456", "bajulaye number 1", "password")
	if koloBank == nil {
		t.Error("error creating kolo bank")
	}
	if koloBank.countNumberOfAccounts() != 1 {
		t.Error("Expected number of accounts to be 1")
	}
	err := koloBank.Deposit("123456", 2000)

	if err != nil {
		t.Error("error depositing to kolo bank")
	}

	if koloBank.CheckBalance("123456", "password") != 2000 {
		t.Error("Expected balance to be 2000")
	}

	fmt.Println(theCreatedAccount.CheckAccountBalance("123456", "password"))
}
