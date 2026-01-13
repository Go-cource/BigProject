package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comps := SelectAllComps()
		tmpl, _ := template.ParseFiles("../web/templates/index.html")
		tmpl.Execute(w, comps)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasks := SelectAllTasks()
		tmpl, _ := template.ParseFiles("../web/templates/tasks.html")
		tmpl.Execute(w, tasks)
	}
}

func CreateTaskhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comps := SelectAllComps()
		tmpl, _ := template.ParseFiles("../web/templates/createTask.html")
		tmpl.Execute(w, comps)
	}
}
