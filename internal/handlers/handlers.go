package handlers

import (
	"BigProject/internal/db"
	"BigProject/internal/models"
	"encoding/json"
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
	if r.Method == http.MethodPost {
		var newCompTaskResult models.CompTaskResult
		err := json.NewDecoder(r.Body).Decode(&newCompTaskResult)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = db.UpdateCompTask(newCompTaskResult)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(`{"status": "ok"}`))
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
		err := db.InsertNewTask(NewTask)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		http.Redirect(w, r, "/tasks", http.StatusFound)
	}
}

func AskTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")
		allTasks := db.SelectAllTasks()

		var nameCompTasks = make([]models.Task, 0, 1)

		for _, v := range allTasks {
			if v.TaskComp == name {
				nameCompTasks = append(nameCompTasks, v)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(nameCompTasks)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}
