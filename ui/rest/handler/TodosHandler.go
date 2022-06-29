package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"todos.example/biz_core/domain"

	"github.com/gorilla/mux"

	dto_mapper "todos.example/ui/rest/dto-mapper"

	"todos.example/biz_core/usecase"

	"todos.example/ui/rest/model"
)

func TodosPostHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("RestAPI::TodosPostHandler() begin...\n")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var postModel model.TodosPostRequestModel
	json.Unmarshal(reqBody, &postModel)

	fmt.Printf("TodosHandler::TodosPostHandler() request body: %+v\n", postModel)

	var uc = usecase.TaskUC{}
	var createdTask = uc.CreateNewTask(postModel.Username, postModel.TaskDesc)

	var responseModel = dto_mapper.TaskItem2Model(&createdTask)
	json.NewEncoder(rw).Encode(responseModel)
}

func TodosGetHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("TodosHandler::TodosGetHandler() begin...\n")
	var query = r.URL.Query()
	var user = query.Get("user")
	if len(user) > 0 {
		var uc = usecase.TaskUC{}
		var tasks = uc.GetUserTasks(user)
		fmt.Printf("TodosHandler::TodosGetHandler() tasks size: %d\n", len(tasks))

		var responseModels = dto_mapper.TaskList2Model(&tasks)
		json.NewEncoder(rw).Encode(responseModels)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func TodosGetItemHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("RestAPI::TodosGetItemHandler() begin...\n")
	var taskID = mux.Vars(r)["taskID"]

	var uc = usecase.TaskUC{}
	var task = uc.GetTaskByID(taskID)

	if (task != domain.Task{}) {
		var responseModel = dto_mapper.TaskItem2Model(&task)
		json.NewEncoder(rw).Encode(responseModel)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}
