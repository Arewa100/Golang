package utils

import (
	"fmt"
	"strconv"
	"time"
	"toDoList/src/data/models"
	"toDoList/src/dtos/request"
)

func CreateTaskMapper(request request.CreateTaskRequest) models.Task {
	theYear, theMonth, theDay := getDate(request)
	fmt.Println(theYear, theMonth, theDay)
	newTask := models.Task{
		UserId:      request.UserId,
		Title:       request.Title,
		TaskContent: request.TaskContent,
		TaskDate:    time.Date(theYear, theMonth, theDay, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(newTask.TaskDate.Year())
	return newTask
}

func getDate(request request.CreateTaskRequest) (theDay int, theMonth time.Month, theYear int) {
	date := request.TaskDate
	var day, month, year string
	count := 0
	var dates []byte = []byte(date)
	for _, value := range dates {
		if count == 0 {
			if value != '/' {
				day += string(value)
			} else {
				count++
			}
		} else if count == 1 {
			if value != '/' {
				month += string(value)
			} else {
				count++
			}
		} else if count == 2 {
			year += string(value)
		}
	}
	theYear, _ = strconv.Atoi(year)
	currentMonth, _ := strconv.Atoi(month)
	theMonth = time.Month(currentMonth)
	theDay, _ = strconv.Atoi(day)
	return theDay, theMonth, theYear
}
