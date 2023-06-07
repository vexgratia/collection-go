package pool

import (
	"context"
)

type WorkerStatus int

const (
	WORKER_STATUS_NEW WorkerStatus = iota
	WORKER_STATUS_READY
	WORKER_STATUS_IN_PROCESS
	WORKER_STATUS_DONE
	WORKER_STATUS_RECOVERED
	WORKER_STATUS_TERMINATED
	WORKER_STATUS_TERMINATED_TIMEOUT
)

type Worker[T any] struct {
	ID     int
	Pool   *Pool[T]
	Status WorkerStatus
	CTX    context.Context
	Input  chan *Task[T]
	Output chan *Task[T]
}

func NewWorker[T any](id int) *Worker[T] {
	return &Worker[T]{
		ID:     id,
		Status: WORKER_STATUS_NEW,
		Output: make(chan *Task[T]),
	}
}
func (w *Worker[T]) Start() {
	for {
		select {
		case <-w.CTX.Done():
			break
		case task := <-w.Input:
			w.Process(task)
		default:
			continue
		}
	}
}

func (w *Worker[T]) Process(task *Task[T]) error {
	task.Status = TASK_STATUS_IN_PROCESS
	err := task.Process()
	if err != nil {
		task.Status = TASK_STATUS_ERROR
		return err
	}
	task.Status = TASK_STATUS_PROCESSED
	w.Output <- task
	return nil
}
