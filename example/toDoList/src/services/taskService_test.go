package services

import (
	"fmt"
	"testing"
	"toDoList/src/dtos/request"
	"toDoList/src/dtos/response"
)

func TestToCreateATask(test *testing.T) {
	newTaskRepoService := CreateTaskRepoService()
	newTaskRequest := new(request.CreateTaskRequest)
	newTaskRequest.UserId = "Miracle"
	newTaskRequest.Title = "Prayer Request"
	newTaskRequest.TaskContent = "thank you father"
	newTaskRequest.TaskDate = "12/08/2025"
	message := newTaskRepoService.CreateTask(*newTaskRequest)
	if message.Message != "task created successfully" {
		test.Error(message)
	}
}

func TestThatToViewTaskContent(test *testing.T) {
	newTaskRepoService := CreateTaskRepoService()
	newTaskRequest := new(request.CreateTaskRequest)
	newTaskRequest.UserId = "Miracle"
	newTaskRequest.Title = "Prayer Request"
	newTaskRequest.TaskContent = "God is great"
	message := newTaskRepoService.CreateTask(*newTaskRequest)
	if message.Message != "task created successfully" {
		test.Error("task creation error")
	}
	var content response.ViewTaskContentResponse = newTaskRepoService.ViewTaskContent(request.ViewTaskContentRequest{UserId: "Miracle", Title: "Prayer Request"})
	fmt.Println(content.Content)
	if content.Content != "God is great" {
		test.Error("expected content thank you father, got " + content.Content)
	}
}

//func TestToCreateTaskAndDeleteTask(test *testing.T) {
//
//}
