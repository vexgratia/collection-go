package pool

import "context"

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
	Status WorkerStatus
	CTX    context.Context
	Output chan *Task[T]
}

func NewWorker[T any](id int) *Worker[T] {
	return &Worker[T]{
		ID:     id,
		Status: WORKER_STATUS_NEW,
		Output: make(chan *Task[T]),
	}
}
