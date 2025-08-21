package services

import (
	"fmt"
	"toDoList/src/data/repository"
	"toDoList/src/dtos/request"
	"toDoList/src/dtos/response"
	"toDoList/src/utils"
)

type TaskRepoService struct {
	theRepository *repository.TaskRepository
}

func CreateTaskRepoService() *TaskRepoService {
	return &TaskRepoService{
		theRepository: repository.CreateRepository(),
	}
}
func (taskService *TaskRepoService) CreateTask(request request.CreateTaskRequest) response.CreateTaskResponse {
	newTask := utils.CreateTaskMapper(request)
	err := taskService.theRepository.AddTask(&newTask)
	if err != nil {
		newResponse := new(response.CreateTaskResponse)
		newResponse.Message = err.Error()
		return *newResponse
	}
	newResponse := new(response.CreateTaskResponse)
	newResponse.Message = "task created successfully"
	return *newResponse

}

func (taskService *TaskRepoService) ViewTaskContent(request request.ViewTaskContentRequest) response.ViewTaskContentResponse {
	foundedTask, err := taskService.theRepository.GetTask(request.UserId, request.Title)
	if err != nil {
		newResponse := new(response.ViewTaskContentResponse)
		newResponse.Content = err.Error()
		return *newResponse
	}
	if foundedTask == nil {
		newResponse := new(response.ViewTaskContentResponse)
		newResponse.Content = "task not found"
		return *newResponse
	}
	newResponse := new(response.ViewTaskContentResponse)
	newResponse.Content = foundedTask.TaskContent
	fmt.Println(foundedTask.TaskDate)
	newResponse.TaskDate = foundedTask.TaskDate.String()
	return *newResponse
}

func (taskService *TaskRepoService) DeleteTask(request request.DeleteTaskRequest) response.DeleteTaskResponse {
	err := taskService.theRepository.DeleteTask(request.UserId, request.Title)
	if err != nil {
		message := response.DeleteTaskResponse{Message: err.Error()}
		return message
	}
	message := new(response.DeleteTaskResponse)
	message.Message = "task deleted successfully"
	return *message
}
