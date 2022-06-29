package domain

type Task struct {
	ID    string
	Desc  string
	state string
}

type Tasks []Task

func NewTask(id string, desc string) Task {
	return Task{id, desc, "NEW"}
}

func (t Task) IsNewTask() bool {
	return t.state == "NEW"
}

func (t Task) GetState() string {
	return t.state
}
