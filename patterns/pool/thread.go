package pool

import "context"

type ThreadStatus int

const (
	THREAD_STATUS_NEW ThreadStatus = iota
	THREAD_STATUS_READY
	THREAD_STATUS_IN_PROCESS
	THREAD_STATUS_DONE
	THREAD_STATUS_RECOVERED
	THREAD_STATUS_TERMINATED
	THREAD_STATUS_TERMINATED_TIMEOUT
)

type Thread[T any] struct {
	ID     int
	Status ThreadStatus
	Config ThreadConfig[T]
	CTX    context.Context
	Output chan *Task[T]
}
type ThreadConfig[T any] struct {
	Generate func(current uint64, example T) (*Task[T], error)
	Example  T
}

func NewThread[T any](id int, config ThreadConfig[T]) *Thread[T] {
	return &Thread[T]{
		ID:     id,
		Status: THREAD_STATUS_NEW,
		Config: config,
		Output: make(chan *Task[T]),
	}
}
func NewThreadConfig[T any](generate func(current uint64, example T) (*Task[T], error), example T) *ThreadConfig[T] {
	return &ThreadConfig[T]{
		Generate: generate,
		Example:  example,
	}
}
