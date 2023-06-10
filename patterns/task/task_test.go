package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type TestCase[I, O any] struct {
	id     int
	name   string
	input  []I
	output []O
}

var cases = []TestCase[int, int]{
	{0, "empty", []int{}, []int{}},
	{1, "default", []int{1, 2, 3}, []int{4, 5, 6}},
}

func TestTask(t *testing.T) {
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			testIsReady(t, test)
			testDefault(t, test)
			testPanic(t, test)
			testTimeout(t, test)
			testError(t, test)
			testCTX(t, test)
		})
	}
}
func testIsReady[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("ready",
		func(t *testing.T) {
			ctx := context.Background()
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Execute()
			if task.Status != TASK_STATUS_TERMINATED {
				t.Errorf("want %d, got %d", TASK_STATUS_TERMINATED, task.Status)
			}
		})
}
func testDefault[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("default",
		func(t *testing.T) {
			ctx := context.Background()
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Prepare(0,
				func(input []I) ([]O, error) {
					return make([]O, 12), nil
				},
				make([]I, 34))
			if task.Status != TASK_STATUS_READY {
				t.Errorf("want %d, got %d", TASK_STATUS_READY, task.Status)
			}
			task.Execute()
			if task.Status != TASK_STATUS_PROCESSED {
				t.Errorf("want %d, got %d", TASK_STATUS_PROCESSED, task.Status)
			}
		})
}
func testPanic[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("panic",
		func(t *testing.T) {
			ctx := context.Background()
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Prepare(0,
				func(input []I) ([]O, error) {
					input[len(input)] = input[0]
					return make([]O, 12), nil
				},
				make([]I, 34))
			if task.Status != TASK_STATUS_READY {
				t.Errorf("want %d, got %d", TASK_STATUS_READY, task.Status)
			}
			task.Execute()
			if task.Status != TASK_STATUS_RECOVERED {
				t.Errorf("want %d, got %d", TASK_STATUS_RECOVERED, task.Status)
			}
		})
}

func testTimeout[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("timeout",
		func(t *testing.T) {
			ctx := context.Background()
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Prepare(1,
				func(input []I) ([]O, error) {
					time.Sleep(time.Minute)
					return make([]O, 12), nil
				},
				make([]I, 34))
			if task.Status != TASK_STATUS_READY {
				t.Errorf("want %d, got %d", TASK_STATUS_READY, task.Status)
			}
			task.Execute()
			if task.Status != TASK_STATUS_TERMINATED_TIMEOUT {
				t.Errorf("want %d, got %d", TASK_STATUS_TERMINATED_TIMEOUT, task.Status)
			}
		})
}
func testError[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("error",
		func(t *testing.T) {
			ctx := context.Background()
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Prepare(1,
				func(input []I) ([]O, error) {
					return make([]O, 12), fmt.Errorf("test error")
				},
				make([]I, 34))
			if task.Status != TASK_STATUS_READY {
				t.Errorf("want %d, got %d", TASK_STATUS_READY, task.Status)
			}
			task.Execute()
			if task.Status != TASK_STATUS_PROCESSED {
				t.Errorf("want %d, got %d", TASK_STATUS_PROCESSED, task.Status)
			}
			if task.Error == nil {
				t.Errorf("want %v, got %v", task.Error, nil)
			}
		})
}
func testCTX[I, O any](t *testing.T, test TestCase[I, O]) {
	t.Run("context",
		func(t *testing.T) {
			ctx, stop := context.WithCancel(context.Background())
			task := NewTask[I, O](ctx, test.id, test.name)
			if task.Status != TASK_STATUS_NEW {
				t.Errorf("want %d, got %d", TASK_STATUS_NEW, task.Status)
			}
			task.Prepare(1,
				func(input []I) ([]O, error) {
					time.Sleep(time.Minute)
					return make([]O, 12), nil
				},
				make([]I, 34))
			if task.Status != TASK_STATUS_READY {
				t.Errorf("want %d, got %d", TASK_STATUS_READY, task.Status)
			}
			stop()
			task.Execute()
			if task.Status != TASK_STATUS_TERMINATED {
				t.Errorf("want %d, got %d", TASK_STATUS_TERMINATED, task.Status)
			}
		})
}
