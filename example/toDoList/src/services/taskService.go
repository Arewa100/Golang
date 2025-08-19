package services

import (
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
func (theRepository *TaskRepoService) CreateTask(request request.CreateTaskRequest) response.CreateTaskResponse {
	newTask := utils.CreateTaskMapper(request)
	err := theRepository.theRepository.AddTask(&newTask)
	if err != nil {
		newResponse := new(response.CreateTaskResponse)
		newResponse.Message = err.Error()
		return *newResponse
	}
	newResponse := new(response.CreateTaskResponse)
	newResponse.Message = "task created successfully"
	return *newResponse

}
