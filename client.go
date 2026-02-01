package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const addr = "http://localhost:8080"
const CompName = "Comp1"

type Task struct {
	TaskId           int
	TaskComp         string
	TaskCreationTime string
	TaskText         string
	TaskResult       string
	TaskResultTime   string
}

func getMyTasks(path string) []Task {
	tasks := make([]Task, 0, 1)
	myTasks, err := http.Get(addr + path)
	if err != nil {
		fmt.Println(err)
		return tasks
	}
	err = json.NewDecoder(myTasks.Body).Decode(&tasks)
	if err != nil {
		fmt.Println(err)
		return tasks
	}
	return tasks
}

func main() {
	fmt.Println(getMyTasks("/ask_tasks?name=" + CompName))
}
