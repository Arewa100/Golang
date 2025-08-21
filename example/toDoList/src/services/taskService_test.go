package services

import (
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
	newTaskRequest.TaskDate = "10/12/2024"
	message := newTaskRepoService.CreateTask(*newTaskRequest)
	if message.Message != "task created successfully" {
		test.Error("task creation error")
	}
	var content response.ViewTaskContentResponse = newTaskRepoService.ViewTaskContent(request.ViewTaskContentRequest{UserId: "Miracle", Title: "Prayer Request"})
	if content.Content != "God is great" {
		test.Error("expected content thank you father, got " + content.Content)
	}
}

func TestToCreateTaskAndDeleteTask(test *testing.T) {
	newTaskRepoService := CreateTaskRepoService()
	newTaskRequest := new(request.CreateTaskRequest)
	newTaskRequest.UserId = "Miracle"
	newTaskRequest.Title = "Prayer Request"
	newTaskRequest.TaskContent = "God is great"
	newTaskRequest.TaskDate = "10/12/2024"
	message := newTaskRepoService.CreateTask(*newTaskRequest)
	if message.Message != "task created successfully" {
		test.Error("task creation error")
	}

	content := newTaskRepoService.ViewTaskContent(request.ViewTaskContentRequest{UserId: "Miracle", Title: "Prayer Request"})
	if content.Content != "God is great" {
		test.Error("expected content thank you father, got " + content.Content)
	}

	deleteTaskResponse := newTaskRepoService.DeleteTask(request.DeleteTaskRequest{UserId: "Miracle", Title: "Prayer Request"})
	if deleteTaskResponse.Message != "task deleted successfully" {
		test.Error("expected task response to be task deleted successfully")
	}
	expectedContent := newTaskRepoService.ViewTaskContent(request.ViewTaskContentRequest{UserId: "Miracle", Title: "Prayer Request"})
	if expectedContent.Content != "task not found" {
		test.Error("expected: " + expectedContent.Content)
	}

}
