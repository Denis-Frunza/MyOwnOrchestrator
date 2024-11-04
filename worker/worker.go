package worker

import (
	"Cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     *queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I'll collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I'll start task or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I'll start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("I'll stop a task")
}
