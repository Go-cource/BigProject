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
	m.HandleFunc("/create_task", handlers.CreateTaskhandler)
	log.Println("Server started...")
	http.ListenAndServe(":8080", m)

}
