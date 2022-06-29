package repo

import (
	"fmt"

	"todos.example/biz_core/domain"
	"todos.example/biz_core/port"
)

// MemoryRepo implement ITaskRepo
type MemoryRepo struct{}

func init() {
	// ensure interface compliance
	var _ port.ITaskRepo = MemoryRepo{}
}

var memoryMap = make(map[string]domain.Tasks) // map: username -> tasks

func (repo MemoryRepo) CreateTaskForUser(username string, task domain.Task) {
	fmt.Printf("MemoryRepo::CreateTaskForUser(): username: %s, task: %+v\n", username, task)
	var _, exist = memoryMap[username]
	if !exist {
		memoryMap[username] = make(domain.Tasks, 0)
	}
	memoryMap[username] = append(memoryMap[username], task)
}

func (repo MemoryRepo) GetTasksOfUser(username string) domain.Tasks {
	return memoryMap[username]
}

func (repo MemoryRepo) GetTaskByID(taskID string) domain.Task {
	for _, tasks := range memoryMap {
		for _, task := range tasks {
			if task.ID == taskID {
				return task
			}
		}
	}
	return domain.Task{}
}

func Dummy() {
	for user, tasks := range memoryMap {
		fmt.Printf("MemoryRepo::Dummy() all tasks of user %s is:\n", user)
		for _, task := range tasks {
			fmt.Printf("MemoryRepo::Dummy()\t - %+v\n", task)
		}
	}
}
