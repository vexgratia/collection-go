package task

import (
	"context"
	"fmt"
	"time"
)

type TaskStatus int

const (
	TASK_STATUS_NEW TaskStatus = iota
	TASK_STATUS_READY
	TASK_STATUS_IN_PROCESS
	TASK_STATUS_PROCESSED
	TASK_STATUS_RECOVERED
	TASK_STATUS_TERMINATED
	TASK_STATUS_TERMINATED_TIMEOUT
)

type Task[I, O any] struct {
	ID       int
	Name     string
	Status   TaskStatus
	CTX      context.Context
	LocalCTX context.Context
	Done     chan struct{}
	Stop     context.CancelFunc
	Error    error
	Timeout  time.Duration
	Process  func(input []I) ([]O, error)
	Input    []I
	Output   []O
}

func NewTask[I, O any](ctx context.Context, id int, name string) *Task[I, O] {
	return &Task[I, O]{
		ID:     id,
		Name:   name,
		Status: TASK_STATUS_NEW,
		CTX:    ctx,
		Done:   make(chan struct{}),
	}
}

func (t *Task[I, O]) Prepare(timeout time.Duration, process func(input []I) ([]O, error), input []I) {
	t.Process = process
	t.Input = input
	t.Timeout = timeout
	if t.Timeout > 0 {
		t.LocalCTX, t.Stop = context.WithTimeout(context.Background(), t.Timeout)
	} else {
		t.LocalCTX, t.Stop = context.WithCancel(context.Background())
	}
	t.Status = TASK_STATUS_READY
}
func (t *Task[I, O]) Ready() bool {
	if t.CTX == nil || t.LocalCTX == nil || t.Done == nil || t.Stop == nil {
		return false
	}
	t.Status = TASK_STATUS_READY
	return true
}

func (t *Task[I, O]) Execute() error {
	if !t.Ready() {
		t.Status = TASK_STATUS_TERMINATED
		return fmt.Errorf("task is not ready")
	}
	t.Status = TASK_STATUS_IN_PROCESS
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Status = TASK_STATUS_RECOVERED
				t.Error = fmt.Errorf("%v", r)
			} else {
				t.Status = TASK_STATUS_PROCESSED
			}
			t.Done <- struct{}{}
		}()
		t.Output, t.Error = t.Process(t.Input)
	}()
	for {
		select {
		case <-t.Done:
			if t.Error != nil {
				return t.Error
			}
			return nil
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
