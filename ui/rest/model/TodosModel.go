package model

type TodosPostRequestModel struct {
	Username string `json:"userName"`
	TaskDesc string `json:"taskDesc"`
}

type TodosPostResponseModel struct {
	TaskID string          `json:"taskid"`
	Detail TodosItemDetail `json:"detail"`
}

type TodosItemDetail struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}
