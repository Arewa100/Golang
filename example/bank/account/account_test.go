package account

import (
	"errors"
	"testing"
)

func createAccount() *Account {
	var newAccount *Account = CreateAccount("Ola", "123456", "bajulaye number 1", "password")
	return newAccount
}

func TestToCreateAnAccount_AccountIsReturned(t *testing.T) {
	var newAccount *Account = createAccount()
	validateAccountCreation(newAccount, t)
	if newAccount != nil && newAccount.name == "Ola" {
		t.Log("account created successfully")
	}
}

func TestToCheckThatAccountBalanceIsZeroAdAccountCreation(t *testing.T) {
	newAccount := createAccount()
	validateAccountCreation(newAccount, t)
	accountBalance := newAccount.CheckAccountBalance("123445", "password")
	if accountBalance != 0.0 {
		t.Error("account balance error")
	}
}

func TestToDeposit1000_BalanceIs1000(t *testing.T) {
	newAccount := createAccount()
	validateAccountCreation(newAccount, t)
	if newAccount.CheckAccountBalance("123456", "password") != 0.0 {
		t.Error("account balance not equal to zero")
	}
	err := newAccount.Deposit("123456", 1000)
	if err != nil {
		t.Error(err)
	}
	if newAccount.CheckAccountBalance("123456", "password") != 1000 {
		t.Error("account balance error")
	}

}

func validateAccountCreation(account *Account, t *testing.T) {
	if account == nil {
		t.Error("error creating new account")
	}
}

func TestThatToDepositA_NegativeAmount_Error_IsThrown(t *testing.T) {
	newAccount := createAccount()
	err := newAccount.Deposit("123456", -1)
	expectedErr := errors.New("account number is not valid")
	if errors.Is(expectedErr, err) {
		t.Error(err)
	}
}
