package pool

import (
	"context"

	. "github.com/vexgratia/collection-go/patterns/task"
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

type Worker[I, O any] struct {
	ID     int
	Pool   *Pool[I, O]
	Status WorkerStatus
	CTX    context.Context
	Input  chan *Task[I, O]
	Output chan *Task[I, O]
}

func NewWorker[I, O any](id int) *Worker[I, O] {
	return &Worker[I, O]{
		ID:     id,
		Status: WORKER_STATUS_NEW,
		Output: make(chan *Task[I, O]),
	}
}
