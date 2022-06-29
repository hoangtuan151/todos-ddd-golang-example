package port

import "todos.example/biz_core/domain"

type ITaskRepo interface {
	CreateTaskForUser(username string, task domain.Task)
	GetTasksOfUser(username string) domain.Tasks
	GetTaskByID(taskID string) domain.Task
}
