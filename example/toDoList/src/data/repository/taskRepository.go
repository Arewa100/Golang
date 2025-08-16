package repository

import (
	"errors"
	"toDoList/src/data/models"
)

type TaskRepository struct {
	UserTaskRepository map[int]models.Task
}

func CreateRepository() *TaskRepository {
	return &TaskRepository{
		UserTaskRepository: make(map[int]models.Task),
	}
}

func (taskRepo *TaskRepository) AddTask(task *models.Task) error {
	if &task != nil {
		currentLength := len(taskRepo.UserTaskRepository)
		taskRepo.UserTaskRepository[currentLength+1] = *task
		return nil
	}
	return nil
}

func (taskRepo *TaskRepository) GetTask(userId string, title string) (task *models.Task, err error) {
	count := len(taskRepo.UserTaskRepository)
	if count == 0 {
		err = errors.New("task repository is empty")
	}
	for _, eachTask := range taskRepo.UserTaskRepository {
		if eachTask.UserId == userId && eachTask.Title == title {
			task = &eachTask
			return task, nil
		}
	}
	return nil, errors.New("task not found")

}

func (taskRepo *TaskRepository) Count() float64 {
	return float64(len(taskRepo.UserTaskRepository))
}
