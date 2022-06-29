package main

import (
	"fmt"

	"todos.example/ui/rest"
)

func init() {
	fmt.Println("main::init()...")
}

func main() {
	fmt.Println("begin...")
	var restApi = rest.RestAPI{Port: "9669"}
	restApi.StartListen()
}
