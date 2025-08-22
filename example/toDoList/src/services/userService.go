package services

import (
	"toDoList/src/data/models"
	"toDoList/src/data/repository"
	"toDoList/src/dtos/request"
	"toDoList/src/dtos/response"
)

type UserService struct {
	userServiceRepository *repository.UserRepository
	TaskRepoService       *TaskRepoService
}

func CreateUserService() *UserService {
	return &UserService{
		userServiceRepository: repository.CreateUserRepository(),
		TaskRepoService:       CreateTaskRepoService(),
	}
}
func (userService *UserService) register(request request.RegisterUserRequest) response.RegisterUserResponse {
	newUser := models.User{
		UserName: request.Username,
		Password: request.Password,
	}
	err := userService.userServiceRepository.AddUser(&newUser)
	if err != nil {
		return response.RegisterUserResponse{Message: err.Error()}
	}
	return response.RegisterUserResponse{Message: "user registered successfully"}
}

func (userService *UserService) AddTask(request request.CreateTaskRequest) response.CreateTaskResponse {
	theUser, err := userService.userServiceRepository.FindUserByUserName(request.UserId)
	if err != nil {
		return response.CreateTaskResponse{Message: err.Error()}
	}
	if theUser == nil {
		return response.CreateTaskResponse{Message: "user not found"}
	}
	if theUser.IsLoggedIn == false {
		return response.CreateTaskResponse{Message: "user is not logged in"}
	} else {
		createTaskResponse := userService.TaskRepoService.CreateTask(request)
		return response.CreateTaskResponse{Message: createTaskResponse.Message}
	}
}

func (userService *UserService) LoginUser(request request.LoginUserRequest) response.LoginUserResponse {
	theUser, err := userService.userServiceRepository.FindUserByUserName(request.Username)
	if err != nil {
		return response.LoginUserResponse{Message: err.Error()}
	}
	if theUser == nil {
		return response.LoginUserResponse{Message: "user not found"}
	}
	if theUser.IsLoggedIn == false {
		theUser.IsLoggedIn = true
		return response.LoginUserResponse{Message: "user logged in successfully"}
	}
	return response.LoginUserResponse{Message: "user already logged in"}
}

func (userService *UserService) LogoutUser(request request.LogoutUserRequest) response.LogoutUserResponse {
	theUser, err := userService.userServiceRepository.FindUserByUserName(request.UserName)
	if err != nil {
		return response.LogoutUserResponse{Message: err.Error()}
	}
	if theUser == nil {
		return response.LogoutUserResponse{Message: "user not found"}
	}
	if theUser.IsLoggedIn == true {
		theUser.IsLoggedIn = false
		return response.LogoutUserResponse{Message: "user logged out successfully"}
	}
	return response.LogoutUserResponse{Message: "logout error: check username"}
}
