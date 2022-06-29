package dto_mapper

import (
	"todos.example/biz_core/domain"
	"todos.example/ui/rest/model"
)

func TaskItem2Model(task *domain.Task) *model.TodosPostResponseModel {
	return &model.TodosPostResponseModel{
		TaskID: task.ID,
		Detail: model.TodosItemDetail{
			Description: task.Desc,
			Status:      task.GetState(),
		},
	}
}

func TaskList2Model(tasks *domain.Tasks) *[]model.TodosPostResponseModel {
	var result = make([]model.TodosPostResponseModel, len(*tasks))
	for idx, task := range *tasks {
		result[idx] = *TaskItem2Model(&task)
	}
	return &result
}
