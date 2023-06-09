package queue

import (
	"testing"

	. "github.com/vexgratia/collection-go/helpers/test"
)

type TestCase[T any] struct {
	name  string
	queue *Queue[T]
	enq   T
}

func TestQueue(t *testing.T) {
	t.Run("integer",
		func(t *testing.T) {
			cases := []TestCase[int]{
				{"empty", NewQueue[int](), 123},
				{"default", &Queue[int]{[]int{0, 123, 456}}, -789},
			}
			for _, test := range cases {
				testBehaviour(t, &test)
			}
		})
	t.Run("string",
		func(t *testing.T) {
			cases := []TestCase[string]{
				{"empty", NewQueue[string](), "abc"},
				{"default", &Queue[string]{[]string{"def", "123", "!@#"}}, "(())"},
			}
			for _, test := range cases {
				testBehaviour(t, &test)
			}
		})
	t.Run("float64",
		func(t *testing.T) {
			cases := []TestCase[float64]{
				{"empty", NewQueue[float64](), 1.23},
				{"default", &Queue[float64]{[]float64{4.56, 7.89, -1.23}}, 9.99},
			}
			for _, test := range cases {
				testBehaviour(t, &test)
			}
		})
}
func testBehaviour[T any](t *testing.T, test *TestCase[T]) {
	t.Run(test.name, func(t *testing.T) {
		testLen(t, test)
		testIsEmpty(t, test)
		testCopy(t, test)
		testPeek(t, test)
		testEnqueue(t, test)
		testDequeue(t, test)
		testClear(t, test)
	})
}
func testLen[T any](t *testing.T, test *TestCase[T]) {
	t.Run("len",
		func(t *testing.T) {
			AssertNotPanic(t, func() {
				length := len(test.queue.Data)
				AssertEqual(t, test.queue.Len(), length)
			})
		})
}
func testIsEmpty[T any](t *testing.T, test *TestCase[T]) {
	t.Run("empty",
		func(t *testing.T) {
			AssertNotPanic(t, func() {
				length := len(test.queue.Data)
				switch length {
				case 0:
					AssertTrue(t, test.queue.Empty())
				default:
					AssertFalse(t, test.queue.Empty())
				}
			})
		})
}
func testCopy[T any](t *testing.T, test *TestCase[T]) {
	t.Run("copy",
		func(t *testing.T) {
			copy := test.queue.Copy()
			AssertEqual(t, *copy, *test.queue)
		})
}
func testPeek[T any](t *testing.T, test *TestCase[T]) {
	t.Run("peek",
		func(t *testing.T) {
			length := len(test.queue.Data)
			switch length {
			case 0:
				AssertPanic(t, func() { test.queue.Peek() })
			default:
				AssertNotPanic(t, func() { test.queue.Peek() })
				peek := test.queue.Peek()
				AssertEqual(t, peek, test.queue.Data[0])
			}
		})
}
func testEnqueue[T any](t *testing.T, test *TestCase[T]) {
	t.Run("enqueue",
		func(t *testing.T) {
			AssertNotPanic(t, func() {
				test.queue.Enqueue(test.enq)
				length := len(test.queue.Data)
				AssertEqual(t, test.queue.Data[length-1], test.enq)
			})
		})
}
func testDequeue[T any](t *testing.T, test *TestCase[T]) {
	t.Run("dequeue",
		func(t *testing.T) {
			length := len(test.queue.Data)
			switch length {
			case 0:
				AssertPanic(t, func() { test.queue.Dequeue() })
			default:
				AssertNotPanic(t, func() {
					want := test.queue.Data[0]
					deq := test.queue.Dequeue()
					AssertEqual(t, want, deq)
				})
			}
		})
}
func testClear[T any](t *testing.T, test *TestCase[T]) {
	t.Run("clear",
		func(t *testing.T) {
			AssertNotPanic(t, func() {
				want := NewQueue[T]()
				test.queue.Clear()
				AssertEqual(t, *test.queue, *want)
			})
		})
}
