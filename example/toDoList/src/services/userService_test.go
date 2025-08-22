package services

import (
	"testing"
	"toDoList/src/dtos/request"
)

func TestThatToCreateAUserService(test *testing.T) {
	newUserService := CreateUserService()
	if newUserService.userServiceRepository == nil {
		test.Errorf("user service respository is nil")
	}
}

func TestToRegisterAUser(test *testing.T) {
	newUserService := CreateUserService()
	if newUserService.userServiceRepository == nil {
		test.Errorf("user service repository is nil")
	}
	registerUserRequest := request.RegisterUserRequest{Username: "Olasoyin", Password: "123456"}
	response := newUserService.register(registerUserRequest)
	if response.Message != "user registered successfully" {
		test.Errorf("expected message to be \"user registered successfully\", got %s", response.Message)
	}
}

func TestForUserToAddTaskWhenLoggedOut(test *testing.T) {
	newUserService := CreateUserService()
	if newUserService.userServiceRepository == nil {
		test.Errorf("user service repository is nil")
	}
	message := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to martket", TaskDate: "23/08/2025"})
	if message.Message == 
	if message.Message != "user not logged in" {
		test.Error("expected user not logged in")
	}

}
