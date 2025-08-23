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

func (userService *UserService) ViewTask(request request.ViewTaskContentRequest) response.ViewTaskResponse {
	foundedUser, err := userService.userServiceRepository.FindUserByUserName(request.UserId)
	if foundedUser != nil && checkIfUserIsLoggedIn(*foundedUser) == false {
		return response.ViewTaskResponse{Content: "user is not logged in"}
	}
	if err != nil {
		return response.ViewTaskResponse{Content: err.Error()}
	}
	if foundedUser == nil {
		return response.ViewTaskResponse{Content: "user not found"}
	}
	foundedTaskResponse := userService.TaskRepoService.ViewTaskContent(request)
	return response.ViewTaskResponse{Content: foundedTaskResponse.Content, TaskDate: foundedTaskResponse.TaskDate}
}

func checkIfUserIsLoggedIn(user models.User) bool {
	if user.IsLoggedIn == false {
		return false
	}
	return true
}

func (userService *UserService) DeleteTask(request request.DeleteTaskRequest) response.DeleteTaskResponse {
	foundedUser, err := userService.userServiceRepository.FindUserByUserName(request.UserId)
	if foundedUser != nil && checkIfUserIsLoggedIn(*foundedUser) == false {
		return response.DeleteTaskResponse{Message: "user is not logged in"}
	}
	if err != nil {
		return response.DeleteTaskResponse{Message: err.Error()}
	}
	deleteTaskResponse := userService.TaskRepoService.DeleteTask(request)
	return response.DeleteTaskResponse{Message: deleteTaskResponse.Message}
}

func (userService *UserService) UpdateTask(request request.UpdateTaskContentRequest) response.UpdateTaskContentResponse {
	foundedUser, err := userService.userServiceRepository.FindUserByUserName(request.UserId)
	if foundedUser != nil && checkIfUserIsLoggedIn(*foundedUser) == false {
		return response.UpdateTaskContentResponse{Message: "user is not logged in"}
	}
	if err != nil {
		return response.UpdateTaskContentResponse{Message: err.Error()}
	}
	updateResponse := userService.TaskRepoService.UpdateTaskContent(request)
	return response.UpdateTaskContentResponse{Message: updateResponse.Message}
}
