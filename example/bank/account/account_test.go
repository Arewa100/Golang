package account

import (
	"testing"
)

func TestToCreateAnAccount_AccountIsReturned(t *testing.T) {
	var newAccount *Account = CreatAccount("Ola", "123456", "bajulaye number 1")
	if newAccount == nil {
		t.Error("error creating new account")
	}

	if newAccount != nil && newAccount.name == "Ola" {
		t.Log("account created successfully")
	}
}
