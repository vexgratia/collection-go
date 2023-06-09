package task

import (
	"context"
	"time"
)

type TaskStatus int

const (
	TASK_STATUS_NEW TaskStatus = iota
	TASK_STATUS_READY
	TASK_STATUS_IN_PROCESS
	TASK_STATUS_PROCESSED
	TASK_STATUS_ERROR
	TASK_STATUS_RECOVERED
	TASK_STATUS_TERMINATED
	TASK_STATUS_TERMINATED_TIMEOUT
)

type Task[T any] struct {
	ID       int
	Name     string
	Status   TaskStatus
	CTX      context.Context
	LocalCTX context.Context
	Done     chan struct{}
	Stop     context.CancelFunc
	Error    chan error
	Timeout  time.Duration
	Process  func(task *Task[T])
	Data     T
}

func NewTask[T any](id int, name string) *Task[T] {
	return &Task[T]{
		ID:     id,
		Name:   name,
		Status: TASK_STATUS_NEW,
	}
}

func (t *Task[T]) Prepare(timeout time.Duration, process func(task *Task[T]), data T) {
	t.Process = process
	t.Data = data
	t.Status = TASK_STATUS_READY
	if t.Timeout > 0 {
		t.LocalCTX, t.Stop = context.WithTimeout(context.Background(), t.Timeout)
	} else {
		t.LocalCTX, t.Stop = context.WithCancel(context.Background())
	}
}

func (t *Task[T]) Execute() error {
	t.Status = TASK_STATUS_IN_PROCESS
	go func() {
		defer func() {
			t.Done <- struct{}{}
		}()
		defer func() {
			if r := recover(); r != nil {
				t.Status = TASK_STATUS_RECOVERED
			}
		}()
		t.Process(t)
	}()
	for {
		select {
		case <-t.Done:
			t.Status = TASK_STATUS_PROCESSED
			return nil
		case err := <-t.Error:
			t.Status = TASK_STATUS_ERROR
			return err
		case <-t.CTX.Done():
			t.Status = TASK_STATUS_TERMINATED
			return nil
		case <-t.LocalCTX.Done():
			t.Status = TASK_STATUS_TERMINATED_TIMEOUT
			return nil
		default:
			continue
		}
	}
}
