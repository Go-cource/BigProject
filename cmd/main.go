package main

import (
	"BigProject/internal/handlers"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlers.IndexHandler)
	m.HandleFunc("/tasks", handlers.TasksHandler)

	//http://localhost:8080/ask_tasks?name=Comp1
	m.HandleFunc("/ask_tasks", handlers.AskTasksHandler)

	m.HandleFunc("/create_task", handlers.CreateTaskhandler)
	log.Println("Server started...")
	http.ListenAndServe(":8080", m)

}
