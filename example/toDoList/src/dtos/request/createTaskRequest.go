package request

type CreateTaskRequest struct {
	UserId      string
	Title       string
	TaskContent string
	TaskDate    string //12/08/2025
}
