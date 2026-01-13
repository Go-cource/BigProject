package handlers

import (
	"BigProject/internal/db"
	"BigProject/internal/models"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comps := db.SelectAllComps()
		tmpl, _ := template.ParseFiles("./web/templates/index.html")
		tmpl.Execute(w, comps)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasks := db.SelectAllTasks()
		tmpl, _ := template.ParseFiles("./web/templates/tasks.html")
		tmpl.Execute(w, tasks)
	}
}

func CreateTaskhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comps := db.SelectAllComps()
		tmpl, _ := template.ParseFiles("./web/templates/createTask.html")
		tmpl.Execute(w, comps)
	}
	if r.Method == http.MethodPost {
		var NewTask models.Task
		NewTask.TaskComp = r.FormValue("CompName")
		NewTask.TaskText = r.FormValue("TaskText")
		NewTask.TaskCreationTime = time.Now().Format("2006/01/02 15:04:05")
		fmt.Printf("New Task: \n %+v", NewTask)
	}
}
