package models

type Comp struct {
	CompId              int
	CompName            string
	CompLastTimeMessage string
}

type Task struct {
	TaskId           int
	TaskComp         string
	TaskCreationTime string
	TaskText         string
	TaskResult       string
	TaskResultTime   string
}

type CompTaskResult struct {
	TaskId     int
	TaskResult string
}
