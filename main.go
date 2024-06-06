package main

import (
	httpRequest "net/http"

	"github.com/kalelc/hexagonal-arquitecture/application/service"
	"github.com/kalelc/hexagonal-arquitecture/infraestructure/http"
	"github.com/kalelc/hexagonal-arquitecture/infraestructure/persistence"
)

func main() {
	taskRepository := persistence.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	taskHandler := http.NewTaskHandler(taskService)

	httpRequest.HandleFunc("/tasks", taskHandler.GetTasks)
	httpRequest.HandleFunc("/tasks/create", taskHandler.CreateTask)
	httpRequest.ListenAndServe(":8080", nil)
}
