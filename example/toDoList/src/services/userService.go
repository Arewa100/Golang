package services

import (
	"toDoList/src/data/models"
	"toDoList/src/data/repository"
	"toDoList/src/dtos/request"
	"toDoList/src/dtos/response"
)

type UserService struct {
	userServiceRepository *repository.UserRepository
}

func CreateUserService() *UserService {
	return &UserService{
		userServiceRepository: repository.CreateUserRepository(),
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
