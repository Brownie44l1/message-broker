package main

import (
	"errors"
	"fmt"
)

type Task struct {
	id          int
	description string
	status      string
}

type TaskQueue struct {
	taskQueue []Task
	nextID    int
}

func (t *TaskQueue) AddTask(description string) {
	t.nextID++
	task := Task{
		id:          t.nextID,
		description: description,
		status:      "pending",
	}
	t.taskQueue = append(t.taskQueue, task)
}

func (t *TaskQueue) ClaimTask() (*Task, error) {
	for i := range t.taskQueue {
		if t.taskQueue[i].status == "pending" {
			t.taskQueue[i].status = "in_progress"
			return &t.taskQueue[i], nil
		}
	}
	return nil, errors.New("failed to claim task")
}

func (t *TaskQueue) CompleteTask(id int) error {
	for i := range t.taskQueue {
		if id == t.taskQueue[i].id {
			t.taskQueue[i].status = "complete"
			fmt.Printf("Task with ID %d has been %s\n", id, t.taskQueue[i].status)
			return nil
		}
	}
	return errors.New("failed to complete task")
}

func (t *TaskQueue) ListPending() []*Task {
	var pending []*Task
	for i := range t.taskQueue {
		if t.taskQueue[i].status == "pending" {
			pending = append(pending, &t.taskQueue[i])
		}
	}
	return pending
}

func main() {
	taskQueue := &TaskQueue{}

	taskQueue.AddTask("Learn Go")
	taskQueue.AddTask("Build projects")
	taskQueue.AddTask("Become a backend engineer")

	task, err := taskQueue.ClaimTask()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Task: ", task)
	}

	err = taskQueue.CompleteTask(1)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	tasks := taskQueue.ListPending()
	for index, value := range tasks {
		fmt.Printf("Task %d -> %s\n", index, value.description)
	}
}
