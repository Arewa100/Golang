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

func TestToAddUserAndDeleteTheUser(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository == nil {
		test.Error("repo not created successfully")
	}
	newUser := createUserOne()
	if newUser == nil && newUserRepository != nil {
		err := newUserRepository.AddUser(newUser)
		if err != nil {
			test.Error(err.Error())
		}
		if newUserRepository.Count() != 1 {
			test.Error("user not added in repo successfully, expected 1")
		}

		deleteErr := newUserRepository.DeleteUser("arewaking")
		if deleteErr != nil {
			test.Error(deleteErr.Error())
		}
		if newUserRepository.Count() != 0 {
			test.Error("user not deleted in repo successfully, expected 0")
		}
	}

}

func TestToAddUserAndFindUserByUserName(test *testing.T) {
	newUserRepository := CreateUserRepository()
	if newUserRepository == nil {
		test.Error("repo not created successfully")
	}
	newUser := createUserOne()
	if newUser == nil && newUserRepository != nil {
		err := newUserRepository.AddUser(newUser)
		if err != nil {
			test.Error(err.Error())
		}

		if newUserRepository.Count() != 1 {
			test.Error("user not added in repo successfully, expected 1")
		}

		//finding the user
		foundedUser, err := newUserRepository.FindUserByUserName("arewaking")
		if err != nil {
			test.Error("user not found")
		}
		if foundedUser != nil && foundedUser.UserName != "arewaking" {
			test.Error("founded user error")
		}
	}

}

func TestThatAllFieldOfUserMustBeFilled(test *testing.T) {
	newUserRepository := CreateUserRepository()
	newUser := models.User{
		UserName: "Miracle",
		Password: "",
	}
	err := newUserRepository.AddUser(&newUser)
	if err != nil && err.Error() != "password is empty" {
		test.Error("expected password is empty")
	}
}
