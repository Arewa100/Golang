package repository

import (
	"errors"
	"strings"
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
	err := checkIfTaskFieldIsEmpty(task)
	if err != nil {
		return err
	}
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

func (taskRepo *TaskRepository) DeleteTask(userId string, title string) error {
	repoError := checkIfRepoIsEmpty(taskRepo)
	if repoError != nil {
		return repoError
	}
	err := deleteTask(taskRepo, userId, title)
	if err != nil {
		return err
	}
	return nil
}

func deleteTask(taskRepo *TaskRepository, userId string, title string) error {
	for key, eachTask := range taskRepo.UserTaskRepository {
		if eachTask.UserId == userId && eachTask.Title == title {
			delete(taskRepo.UserTaskRepository, key)
			return nil
		}
	}
	return errors.New("task not found")
}

func checkIfRepoIsEmpty(taskRepo *TaskRepository) error {
	count := len(taskRepo.UserTaskRepository)
	if count == 0 {
		return errors.New("task repository is empty")
	}
	return nil
}

func (taskRepo *TaskRepository) UpdateTask(userId string, title string, content string) error {
	for key, eachTask := range taskRepo.UserTaskRepository {
		if eachTask.UserId == userId && eachTask.Title == title {
			newTask := new(models.Task)
			newTask.UserId = eachTask.UserId
			newTask.Title = eachTask.Title
			newTask.TaskDate = eachTask.TaskDate
			newTask.TaskContent = content
			delete(taskRepo.UserTaskRepository, key)
			taskRepo.UserTaskRepository[key] = *newTask
			return nil
		}
	}
	return errors.New("task not found and not updated successfully")
}
func checkIfTaskFieldIsEmpty(task *models.Task) error {
	if strings.TrimSpace(task.UserId) == "" {
		return errors.New("user id is empty")
	}
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("task title is empty")
	}
	if strings.TrimSpace(task.TaskContent) == "" {
		return errors.New("task content is empty")
	}
	if task.TaskDate.IsZero() {
		return errors.New("task date is empty")
	}
	return nil
}
