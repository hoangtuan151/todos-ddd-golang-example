package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"todos.example/ui/rest/handler"
)

type RestAPI struct {
	Port string
}

const (
	apiBasePath       = "/api"
	todosResourcePath = apiBasePath + "/todos"
)

func (api RestAPI) StartListen() {
	var routers = mux.NewRouter()
	routers.Use(commonMiddleware)

	fmt.Printf("RestAPI::StartListen() configuring endpoints...\n")
	routers.HandleFunc(todosResourcePath, handler.TodosPostHandler).Methods(http.MethodPost)
	routers.HandleFunc(todosResourcePath, handler.TodosGetHandler).Methods(http.MethodGet)
	routers.HandleFunc(todosResourcePath+"/{taskID}", handler.TodosGetItemHandler).Methods(http.MethodGet)

	fmt.Printf("RestAPI::StartListen() start listening at port %s\n", api.Port)
	log.Fatal(http.ListenAndServe(":"+api.Port, routers))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(rw, r)
		})
}
