package pool

type ProcessFunc func()

type TaskState int

const (
	TASK_STATE_NEW TaskState = iota
	TASK_STATE_READY
	TASK_STATE_IN_PROCESS
	TASK_STATE_PROCESSED
	TASK_STATE_RECOVERED
	TASK_STATE_TERMINATED
	TASK_STATE_TERMINATED_TIMEOUT
)

type Task struct {
	ID      int
	Name    string
	Process ProcessFunc
	State   TaskState
	Data    any
}
