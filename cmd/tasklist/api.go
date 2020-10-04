package main

import (
	"encoding/json"
	"fmt"
	"go-tutorials/internal/task"
)

func getTasks() ([]byte, error) {
	b, err := json.Marshal(TaskList)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return b, nil
}

func addTask(body []byte) []byte {
	task := task.Task{Status: task.NEW}
	json.Unmarshal(body, &task)
	TaskList.ADD(task.Name, &task)
	return []byte("Task Added")
}
