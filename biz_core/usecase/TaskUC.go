package usecase

import (
	"fmt"

	"todos.example/infra/repo"

	"github.com/google/uuid"
	"todos.example/biz_core/domain"
	"todos.example/biz_core/port"
)

var taskRepository port.ITaskRepo

func init() {
	fmt.Println("TaskUC::init()...")
	taskRepository = repo.MemoryRepo{}
}

type TaskUC struct {
}

func (t TaskUC) CreateNewTask(username string, taskDesc string) domain.Task {
	var taskID = uuid.New().String()
	var task = domain.NewTask(taskID, taskDesc)
	taskRepository.CreateTaskForUser(username, task)
	return task
}

func (t TaskUC) GetUserTasks(username string) domain.Tasks {
	return taskRepository.GetTasksOfUser(username)
}

func (t TaskUC) GetTaskByID(taskID string) domain.Task {
	return taskRepository.GetTaskByID(taskID)
}
