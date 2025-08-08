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
	fmt.Println(theCreatedAccount.CheckAccountBalance("123456", "password"))
	if koloBank.CheckBalance("123456", "password") != 2000 {
		t.Error("Expected balance to be 2000")
	}
}

func TestThatBankCanWithDraw(t *testing.T) {
	koloBank := bank()
	koloBank.CreateAnAccount("Ola", "123456", "bajulaye number 1", "password")
	koloBank.Deposit("123456", 2000)
	if koloBank.CheckBalance("123456", "password") != 2000 {
		t.Error("Expected balance to be 2000")
	}
	koloBank.WithDraw("123456", 1000, "password")
	if koloBank.CheckBalance("123456", "password") != 1000 {
		t.Error("Expected balance to be 2000")
	}
}

func TestThatBankCanTransfer(test *testing.T) {
	koloBank := bank()
	koloBank.CreateAnAccount("Ola", "123456", "bajulaye number 1", "sender_password")
	koloBank.CreateAnAccount("Ola", "144444", "bajulaye number 2", "password")
	koloBank.Deposit("123456", 2000)
	koloBank.Deposit("14444", 2000)
	if koloBank.CheckBalance("144444", "password") != 2000 {
		test.Errorf("Expected balance to be 2000")
	}
	koloBank.Transfer("123456", "144444", 1000, "sender_password")
	if koloBank.CheckBalance("123456", "sender_password") != 1000 {
		test.Errorf("Expected sender balance to be 2000")
	}
	if koloBank.CheckBalance("144444", "password") != 3000 {
		test.Errorf("Expected receiver balance to be 3000")
	}
}
