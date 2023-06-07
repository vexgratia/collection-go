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
	ID        int
	Generator *Generator[T]
	Status    ThreadStatus
	CTX       context.Context
	Output    chan *Task[T]
}

func NewThread[T any](id int) *Thread[T] {
	return &Thread[T]{
		ID:     id,
		Status: THREAD_STATUS_NEW,
		Output: make(chan *Task[T]),
	}
}

func (t *Thread[T]) Start() {
	for {
		select {
		case <-t.CTX.Done():
			break
		default:
			t.Generate()
		}
	}
}
func (t *Thread[T]) Generate() error {
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
