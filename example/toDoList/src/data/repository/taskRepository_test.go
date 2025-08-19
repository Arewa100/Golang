package repository

import (
	"fmt"
	"log"
	"testing"
	"time"
	"toDoList/src/data/models"
)

func createTask() models.Task {
	newTask := models.Task{
		UserId:      "arewaking",
		Title:       "firstTask",
		TaskContent: "i am working on my self",
		TaskDate:    time.Now(),
	}
	return newTask
}

func createSecondTask() models.Task {
	newTask := models.Task{
		UserId:      "arewaking",
		Title:       "secondTask",
		TaskContent: "help me",
		TaskDate:    time.Now(),
	}
	return newTask
}
func createThirdTask() models.Task {
	newTask := models.Task{
		UserId:      "arewaking",
		Title:       "thirdTask",
		TaskContent: "i love you",
		TaskDate:    time.Now(),
	}
	return newTask
}

func TestThatTaskCanBeCreated(t *testing.T) {
	newTaskRepo := CreateRepository()
	if newTaskRepo == nil {
		t.Error("Failed to create task repository")
	}
	if newTaskRepo != nil && newTaskRepo.Count() != 0 {
		t.Error("task repo is not empty")
	}
}

func TestThatTaskRepositoryCanAddTask(t *testing.T) {
	newTaskRepo := CreateRepository()
	if newTaskRepo == nil {
		t.Error("Failed to create task repository")
	}
	newTask := createTask()
	if newTaskRepo != nil {
		err := newTaskRepo.AddTask(&newTask)
		if err != nil {
			t.Error("Failed to add task")
		}
	}

	if newTaskRepo.Count() != 1 {
		t.Error("error adding task, expected 1")
	}
}

func TestToAddMultipleTaskWithSameUserName(test *testing.T) {
	newTaskRepo := CreateRepository()
	firstTask := createTask()
	secondTask := createTask()

	if newTaskRepo != nil {
		firstError := newTaskRepo.AddTask(&firstTask)
		secondError := newTaskRepo.AddTask(&secondTask)
		if newTaskRepo.Count() != 2 {
			test.Error("expected number of task to be 2")
		}

		if firstError != nil || secondError != nil {
			test.Errorf("failed to add task %v\t %v", firstError, secondError)
		}
	}
}

func TestToAddMultipleTaskAndGetTaskByUserName(test *testing.T) {
	newTaskRepo := CreateRepository()
	firstTask := createTask()
	secondTask := createSecondTask()
	thirdTask := createThirdTask()

	if newTaskRepo != nil {
		firstError := newTaskRepo.AddTask(&firstTask)
		secondError := newTaskRepo.AddTask(&secondTask)
		thirdError := newTaskRepo.AddTask(&thirdTask)
		if newTaskRepo.Count() != 3 {
			test.Error("expected number of task to be three 3")
		}
		if firstError != nil || secondError != nil || thirdError != nil {
			test.Errorf("failed to add task %v\t %v %v", firstError, secondError, thirdError)
		}

		task, err := newTaskRepo.GetTask("arewaking", "thirdTask")
		if err != nil {
			test.Error("failed to get task")
		}
		fmt.Println(task.TaskContent)
	}

}

func TestToAddTaskAndDeleteTask_RepoCountIsReduced_From_Two_One(test *testing.T) {
	newTaskRepo := CreateRepository()
	firstTask := createTask()
	secondTask := createSecondTask()

	if newTaskRepo != nil {
		firstError := newTaskRepo.AddTask(&firstTask)
		secondError := newTaskRepo.AddTask(&secondTask)
		if newTaskRepo.Count() != 2 {
			test.Error("expected number of task to be 2")
		}
		if firstError != nil || secondError != nil {
			test.Errorf("failed to add task %v %v", firstError, secondError)
		}

		err := newTaskRepo.DeleteTask("arewaking", "firstTask")
		if err != nil {
			test.Error("task not found")
		}
		if newTaskRepo.Count() != 1 {
			test.Error("error deleting task")
		}
	}
}

func TestToAddTaskAndUpdateTheTaskContent(test *testing.T) {
	newTaskRepo := CreateRepository()
	firstTask := createTask()
	secondTask := createSecondTask()
	if newTaskRepo != nil {
		newError := newTaskRepo.AddTask(&firstTask)
		secondError := newTaskRepo.AddTask(&secondTask)
		if newError != nil || secondError != nil {
			test.Error("failed to add task")
		}

		task, err := newTaskRepo.GetTask("arewaking", "firstTask")
		if err != nil {
			test.Error("failed to get task")
		}

		if task != nil && task.GetTaskContent() != "i am working on my self" {
			test.Error("task content is wrong")
		}

		updateError := newTaskRepo.UpdateTask("arewaking", "firstTask", "i want to start playing football")
		updatedTask, err := newTaskRepo.GetTask("arewaking", "firstTask")
		if updateError != nil {
			test.Error("failed to update task")
			log.Fatal(updateError)
		}
		if updatedTask != nil && updatedTask.GetTaskContent() != "i want to start playing football" {
			fmt.Println(updatedTask.GetTaskContent())
			test.Error("the task not updated successfully")
		}

	}
}

func TestThatAllFieldMustBeFilled(test *testing.T) {
	newTaskRepo := CreateRepository() //thi is where i am 
}
