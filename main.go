package main

import (
	"fmt"

	"todos.example/biz_core/domain"
	"todos.example/biz_core/usecase"
)

func init() {
	fmt.Println("main::init()...")
}

func test1() {
	t := domain.NewTask("123", "my task")
	fmt.Printf("task: %+v\n", t)
	fmt.Printf("task is new: %t\n", t.IsNewTask())

	uc := usecase.TaskUC{}
	uc.CreateNewTask("joe", "first task")
	uc.CreateNewTask("joe", "second task")
	t2 := uc.CreateNewTask("cech", "second task of cech")

	var _ = uc.GetUserTasks("joes")

	t3 := uc.GetTaskByID(t2.ID + "..")
	if (t3 == domain.Task{}) {
		fmt.Println("Not found")
	} else {
		fmt.Printf("Found task: %+v\n", t3)
	}
	// repo.Dummy()
}

func test2() {
	var memory = make(map[string]*[]string)
	var key = "user1"
	var _, exist = memory[key]
	fmt.Printf("is key exist: %t\n", exist)

	if !exist {
		temp := make([]string, 10)
		memory[key] = &temp
		// memory[key] = append(memory[key], "first")
		_ = append(*memory[key], "first")
	}

	fmt.Println(memory)
}

func main() {
	fmt.Println("begin...")
	test1()
}
