package models

import "time"

type Task struct {
	UserId      string
	Title       string
	TaskContent string
	TaskDate    time.Time
}

func (task *Task) GetTaskName() string {
	return task.Title
}

func (task *Task) GetTaskContent() string {
	return task.TaskContent
}

func (task *Task) GetTaskDate() time.Time {
	return task.TaskDate
}

func (task *Task) GetUserId() string {
	return task.UserId
}
