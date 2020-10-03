package task

import "fmt"

const (
	// NEW task
	NEW int = iota
	// DONE is a completed task
	DONE
)

// Task is an element of Tasklist
type Task struct {
	Name   string
	Desc   string
	Status int
}

// Print prints out contents of TData
func (t *Task) Print() {
	fmt.Printf("%+v\n", *t)
}

// GetStatus returns the task status
func (t *Task) GetStatus() int {
	return t.Status
}

// ChangeStatus modifies the task status
func (t *Task) ChangeStatus(s int) {
	t.Status = s
}
