package repository

import "toDoList/src/data/models"

type UserRepository struct {
	UserRepositoryMap map[int]*models.User
}

func CreateUserRepository() *UserRepository {
	return &UserRepository{
		UserRepositoryMap: make(map[int]*models.User),
	}
}

func (userRepository *UserRepository) AddUser(user *models.User) error {
	count := 0
	userRepository.UserRepositoryMap[count+1] = user
	return nil
}

func (userRepository *UserRepository) Count() float32 {
	return float32(len(userRepository.UserRepositoryMap))
}
