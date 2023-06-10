package generator

import (
	"context"

	. "github.com/vexgratia/collection-go/patterns/task"
)

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

type Thread[I, O any] struct {
	ID        int
	Generator *Generator[I, O]
	Status    ThreadStatus
	CTX       context.Context
	Output    chan *Task[I, O]
}

func NewThread[I, O any](id int) *Thread[I, O] {
	return &Thread[I, O]{
		ID:     id,
		Status: THREAD_STATUS_NEW,
		Output: make(chan *Task[I, O]),
	}
}

func (t *Thread[I, O]) Start() {
	for {
		select {
		case <-t.CTX.Done():
			break
		default:
			t.Generate()
		}
	}
}
func (t *Thread[I, O]) Generate() error {
	current := t.Generator.State.Current.Add(1)
	if current >= t.Generator.State.Goal.Load() {
		return nil
	}
	task, err := t.Generator.Generate(current, t.Generator.Example)
	if err != nil {
		return err
	}
	t.Output <- task
	return nil
}
