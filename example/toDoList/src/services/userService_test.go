package services

import (
	"fmt"
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
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	message := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	if message.Message == "user not found" {
		test.Error("user not found: Create user")
	}
	if message.Message != "user is not logged in" {
		test.Error("expected user not logged in")
	}

}
func TestToAddTaskWhenUserIsLoggedIn(test *testing.T) {
	newUserService := CreateUserService()
	if newUserService.userServiceRepository == nil {
		test.Errorf("user service repository is nil")
	}
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	message := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	if message.Message == "user not found" {
		test.Error("user not found: Create user")
	}
	if message.Message != "user is not logged in" {
		test.Error("expected user not logged in")
	}
	loginMessage := newUserService.LoginUser(request.LoginUserRequest{Username: "arewaking", Password: "123456"})
	if loginMessage.Message != "user logged in successfully" {
		fmt.Println(loginMessage.Message)
		test.Error("expected user is logged in")
	}
	addTaskMessage := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	if addTaskMessage.Message != "task created successfully" {
		test.Error("error adding task")
	}

}

func TestToLoginAndLogoutAUser(test *testing.T) {
	newUserService := CreateUserService()
	if newUserService.userServiceRepository == nil {
		test.Errorf("user service repository is nil")
	}
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	loginMessage := newUserService.LoginUser(request.LoginUserRequest{Username: "arewaking", Password: "123456"})
	if loginMessage.Message != "user logged in successfully" {
		fmt.Println(loginMessage.Message)
		test.Error("expected user is logged in")
	}
	addTaskMessage := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	if addTaskMessage.Message != "task created successfully" {
		test.Error("error adding task")
	}
	newUserService.LogoutUser(request.LogoutUserRequest{UserName: "arewaking"})
	message := newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	if message.Message == "user not found" {
		test.Error("user not found: Create user")
	}
	if message.Message != "user is not logged in" {
		test.Error("expected user not logged in")
	}

}

func TestToViewTaskByUser(test *testing.T) {
	newUserService := CreateUserService()
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	newUserService.LoginUser(request.LoginUserRequest{Username: "arewaking", Password: "123456"})
	newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	viewTaskResponse := newUserService.ViewTask(request.ViewTaskContentRequest{UserId: "arewaking", Title: "reminder"})
	if viewTaskResponse.Content != "we are going to market" {
		test.Error("expected we are going to market")
	}

}

func TestToViewContentWhenUserIsNotLoggedIn(test *testing.T) {
	newUserService := CreateUserService()
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	viewTaskResponse := newUserService.ViewTask(request.ViewTaskContentRequest{UserId: "arewaking", Title: "reminder"})
	if viewTaskResponse.Content != "user is not logged in" {
		test.Errorf("expected \"user is not logged in\", but got %s", viewTaskResponse.Content)

	}
}

func TestThatUserCanDeleteTask(test *testing.T) {
	newUserService := CreateUserService()
	newUserService.register(request.RegisterUserRequest{Username: "arewaking", Password: "123456"})
	newUserService.AddTask(request.CreateTaskRequest{UserId: "arewaking", Title: "reminder", TaskContent: "we are going to market", TaskDate: "23/08/2025"})
	viewTaskResponse := newUserService.ViewTask(request.ViewTaskContentRequest{UserId: "arewaking", Title: "reminder"})
	if viewTaskResponse.Content != "user is not logged in" {
		test.Errorf("expected \"user is not logged in\", but got %s", viewTaskResponse.Content)

	}
}
