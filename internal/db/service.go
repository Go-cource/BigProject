package db

import (
	"BigProject/internal/models"
	"database/sql"
	"fmt"
	"log"
)

func SelectAllComps() []models.Comp {

	db := DbConnect()
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM comps`)
	if err != nil {
		log.Println("Query error: ", err)
		return make([]models.Comp, 0)
	}
	defer rows.Close()
	Comps := []models.Comp{}
	for rows.Next() {
		comp := models.Comp{}
		err := rows.Scan(&comp.CompId, &comp.CompName, &comp.CompLastTimeMessage)
		if err != nil {
			fmt.Println("Error with comps scan: ", err)
			continue
		}
		Comps = append(Comps, comp)
	}
	return Comps
}

func SelectAllTasks() []models.Task {

	db := DbConnect()
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		log.Println("Query error: ", err)
		return make([]models.Task, 0)
	}
	defer rows.Close()

	Tasks := []models.Task{}
	for rows.Next() {
		task := models.Task{}
		var TaskResult sql.NullString
		var TaskResultTime sql.NullString
		err := rows.Scan(&task.TaskId, &task.TaskComp, &task.TaskCreationTime, &task.TaskText, &TaskResult, &TaskResultTime)
		if err != nil {
			fmt.Println("Error with tasks scan: ", err)
			continue
		}
		if TaskResult.Valid {
			task.TaskResult = TaskResult.String
		} else {
			task.TaskResult = ""
		}

		if TaskResultTime.Valid {
			task.TaskResultTime = TaskResultTime.String
		} else {
			task.TaskResultTime = ""
		}

		Tasks = append(Tasks, task)
	}
	return Tasks

}
