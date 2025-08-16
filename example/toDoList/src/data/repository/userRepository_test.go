package repository

import (
	"testing"
	"toDoList/src/data/models"
)

func createUserOne() *models.User {
	return &models.User{
		UserName: "Miracle",
		Password: "123456",
	}
}

func TestToCreateUserRepository(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository == nil {
		test.Error("repo not created successfully")
	}
}

func TestThatUserRepository_Can_Add_User(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository == nil && newUserRepository.Count() != 0 {
		test.Error("user repo is not full")
	}
	newUser := createUserOne()
	err := newUserRepository.AddUser(newUser)
	if err != nil {
		test.Error(err.Error())
	}
	if newUserRepository == nil && newUserRepository.Count() != 1 {
		test.Error("user repo is not full")
	}
}
