package repository

import (
	"errors"
	"strings"
	"toDoList/src/data/models"
)

type UserRepository struct {
	UserRepositoryMap map[int]*models.User
}

func CreateUserRepository() *UserRepository {
	return &UserRepository{
		UserRepositoryMap: make(map[int]*models.User),
	}
}

func (userRepository *UserRepository) AddUser(user *models.User) error {
	err := checkAllFields(*user)
	if err != nil {
		return err
	}
	count := 0
	userRepository.UserRepositoryMap[count+1] = user
	return nil
}

func (userRepository *UserRepository) Count() float32 {
	return float32(len(userRepository.UserRepositoryMap))
}

func (userRepository *UserRepository) DeleteUser(userId string) error {
	for key, user := range userRepository.UserRepositoryMap {
		if user.UserName == userId {
			delete(userRepository.UserRepositoryMap, key)
			return nil
		}
	}
	return errors.New("user not found")
}

func (userRepository *UserRepository) FindUserByUserName(userName string) (*models.User, error) {
	for _, user := range userRepository.UserRepositoryMap {
		if user.UserName == userName {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func checkAllFields(user models.User) error {
	if strings.TrimSpace(user.UserName) == "" {
		return errors.New("username is empty")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is empty")
	}
	return nil
}
