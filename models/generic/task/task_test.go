package task

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/vexgratia/collection-go/helpers/test/assert"
)

type TestCase[I, O any] struct {
	inputNewTask   inputNewTask
	inputPrepare   inputPrepare[I, O]
	expectedStatus TaskStatus
	expectedOutput O
	expectedError  error
}
type inputNewTask struct {
	ctx  context.Context
	id   int
	name string
}
type inputPrepare[I, O any] struct {
	timeout time.Duration
	input   I
	process func(input I) (O, error)
}

func TestTask(t *testing.T) {
	var task *Task[int, int]
	var err error
	// for ctx TestCase
	closedCtx, closeCtx := context.WithCancel(context.Background())
	closeCtx()
	//
	var cases = map[string]TestCase[int, int]{
		"empty": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   0,
				name: "empty",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: 0,
				input:   0,
				process: emptyFunc,
			},
			expectedStatus: TASK_STATUS_PROCESSED,
			expectedOutput: 0,
			expectedError:  nil,
		},
		"default": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   1,
				name: "default",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: 0,
				input:   5,
				process: defaultFunc,
			},
			expectedStatus: TASK_STATUS_PROCESSED,
			expectedOutput: 5,
			expectedError:  nil,
		},
		"double": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   2,
				name: "double",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: 0,
				input:   5,
				process: doubleFunc,
			},
			expectedStatus: TASK_STATUS_PROCESSED,
			expectedOutput: 10,
			expectedError:  nil,
		},
		"error": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   3,
				name: "error",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: 0,
				input:   5,
				process: errorFunc,
			},
			expectedStatus: TASK_STATUS_PROCESSED,
			expectedOutput: 0,
			expectedError:  testError,
		},
		"panic": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   4,
				name: "panic",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: 0,
				input:   5,
				process: panicFunc,
			},
			expectedStatus: TASK_STATUS_RECOVERED,
			expectedOutput: 0,
			expectedError:  fmt.Errorf(testPanic),
		},
		"ctx": {
			inputNewTask: inputNewTask{
				ctx:  closedCtx,
				id:   5,
				name: "ctx",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: time.Second,
				input:   5,
				process: timeoutFunc,
			},
			expectedStatus: TASK_STATUS_TERMINATED,
			expectedOutput: 0,
			expectedError:  nil,
		},
		"timeout": {
			inputNewTask: inputNewTask{
				ctx:  context.Background(),
				id:   6,
				name: "timeout",
			},
			inputPrepare: inputPrepare[int, int]{
				timeout: time.Second,
				input:   5,
				process: timeoutFunc,
			},
			expectedStatus: TASK_STATUS_TERMINATED_TIMEOUT,
			expectedOutput: 0,
			expectedError:  nil,
		},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Run("new", func(t *testing.T) {
				// Task creation
				task = NewTask[int, int](test.inputNewTask.ctx, test.inputNewTask.id, test.inputNewTask.name)
				assert.Equal(t, task.ID, test.inputNewTask.id)
				assert.Equal(t, task.Name, test.inputNewTask.name)
				assert.Equal(t, task.Status, TASK_STATUS_NEW)
			})
			t.Run("unprepared", func(t *testing.T) {
				// Task should not be ready yet
				assert.False(t, task.Ready())
				// Trying to execute, should error
				err = task.Execute()
				assert.Equal(t, err, TASK_ERROR[NOT_READY])
			})
			t.Run("prepare", func(t *testing.T) {
				// Task preparation
				task.Prepare(test.inputPrepare.timeout, test.inputPrepare.input, test.inputPrepare.process)
				assert.Equal(t, task.Timeout, test.inputPrepare.timeout)
				assert.Equal(t, task.Input, test.inputPrepare.input)
				assert.Equal(t, reflect.ValueOf(task.Process), reflect.ValueOf(test.inputPrepare.process))
				assert.Equal(t, task.Status, TASK_STATUS_READY)
				// Task should be ready now
				assert.True(t, task.Ready())
			})
			t.Run("execution", func(t *testing.T) {
				// Task execution
				err = task.Execute()
				assert.Equal(t, task.Output, test.expectedOutput)
				assert.Equal(t, task.Error, test.expectedError)
				assert.Equal(t, task.Status, test.expectedStatus)
				assert.Equal(t, err, task.Error)
			})
		})
	}
}
func emptyFunc(x int) (int, error) {
	return 0, nil
}
func defaultFunc(x int) (int, error) {
	return x, nil
}
func doubleFunc(x int) (int, error) {
	return x * 2, nil
}

var testError = fmt.Errorf("error")

func errorFunc(x int) (int, error) {
	return 0, testError
}

var testPanic = fmt.Sprintf("panic")

func panicFunc(x int) (int, error) {
	panic(testPanic)
}
func timeoutFunc(x int) (int, error) {
	time.Sleep(time.Minute)
	return 0, nil
}
