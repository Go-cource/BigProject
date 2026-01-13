package main

import (
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", IndexHandler)
	m.HandleFunc("/tasks", TasksHandler)
	m.HandleFunc("/create_task", CreateTaskhandler)
	log.Println("Server started...")
	http.ListenAndServe(":8080", m)

}
