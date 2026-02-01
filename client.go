package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

const addr = "http://localhost:8080"
const CompName = "Comp1"

type CompTaskResult struct {
	TaskId     int
	TaskResult string
}

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

func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %w", err)
	}
	return string(output), nil
}

func execMyTasks(tasks []Task) error {
	if len(tasks) == 0 {
		return nil
	}
	for _, v := range tasks {
		taskRes := CompTaskResult{}
		taskRes.TaskId = v.TaskId
		tr, err := executeCommand(v.TaskText)
		if err != nil {
			taskRes.TaskResult = "error while executing"
		} else {
			taskRes.TaskResult = tr
		}
		finalRes, err := json.Marshal(taskRes)
		if err != nil {
			fmt.Println(err)
			return err
		}
		postRes, err := http.Post(addr+"/tasks", "application/json", bytes.NewBuffer(finalRes))
		if err != nil {
			fmt.Println(err)
			return err
		}
		srvAnswer := map[string]interface{}{}
		err = json.NewDecoder(postRes.Body).Decode(&srvAnswer)
		if srvAnswer["status"] == "ok" {
			return nil
		}
	}
	return nil
}

func main() {
	// fmt.Println(getMyTasks("/ask_tasks?name=" + CompName))
	err := execMyTasks(getMyTasks("/ask_tasks?name=" + CompName))
	if err != nil {
		fmt.Println(err)
	}
}
