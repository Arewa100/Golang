package services

import (
	"testing"
	"toDoList/src/dtos/request"
)

func TestToCreateATask(test *testing.T) {
	newTaskRepoService := CreateTaskRepoService()
	newTaskRequest := new(request.CreateTaskRequest)
	newTaskRequest.UserId = "Miracle"
	newTaskRequest.Title = "Prayer Request"
	newTaskRequest.TaskContent = "thank you father"
	message := newTaskRepoService.CreateTask(*newTaskRequest)
	if message.Message != "task created successfully" {
		test.Error(message)
	}
}
