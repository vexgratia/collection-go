package pool

type ProcessFunc func()

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

type Task[T any] struct {
	ID      int
	Name    string
	Status  TaskStatus
	Process ProcessFunc
	Data    T
}

func NewTask[T any](id int, name string, process ProcessFunc, data T) *Task[T] {
	return &Task[T]{
		ID:      id,
		Name:    name,
		Process: process,
		Status:  TASK_STATUS_NEW,
		Data:    data,
	}
}
