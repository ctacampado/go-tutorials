package main

import (
	"go-tutorials/internal/kvstore"
	"go-tutorials/internal/task"
)

func main() {
	tList := kvstore.New()

	t := task.Task{Name: "My Task", Desc: "This is a task", Status: task.NEW}
	tList.ADD(t.Name, &t)
	tList.Save()
	tList.PRINT()
	t.Status = task.DONE
	tList.CHANGE(t.Name, &t)
	tList.PRINT()
}
