package main

import (
	"bank/account"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Account Demo")

	newAccount := account.CreateAccount("Ola", "123456", "bajulaye", "password")
	err := newAccount.Deposit("123456", 1000)
	if err != nil {
		fmt.Printf("this error occurs %q", err.Error())
	}
	
}
