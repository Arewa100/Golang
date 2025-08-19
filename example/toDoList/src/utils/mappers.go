package utils

import (
	"time"
	"toDoList/src/data/models"
	"toDoList/src/dtos/request"
)

func CreateTaskMapper(request request.CreateTaskRequest) models.Task {
	newTask := models.Task{
		UserId:      request.UserId,
		Title:       request.Title,
		TaskContent: request.TaskContent,
		TaskDate:    time.Now(),
	}
	return newTask
}
