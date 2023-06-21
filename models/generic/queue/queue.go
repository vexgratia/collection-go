package queue

import (
	"fmt"
	"sync"
)

// A Queue[T] is a generic thread-safe linear data structure.
//
// Queue[T] is based on slice of type T.
//
// Queue[T] is thread-safe.
type Queue[T any] struct {
	mu    *sync.Mutex
	items []T
}

// New creates an empty Queue of type T.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		mu:    &sync.Mutex{},
		items: make([]T, 0),
	}
}

// Len returns current length of the Queue.
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// Empty returns true if the Queue is empty, and false if not.
func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

// Peek returns the last item that's been enqueued.
//
// If the Queue is empty, panic occurs.
func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("can't peek, queue is empty")
	}
	return q.items[0]
}

// PeekAll returns all items from the Queue.
func (q *Queue[T]) PeekAll() []T {
	return q.items
}

// Enqueue inserts items in the end of the Queue.
func (q *Queue[T]) Enqueue(items ...T) {
	q.mu.Lock()
	for _, item := range items {
		q.items = append(q.items, item)
	}
	q.mu.Unlock()
}

// Dequeue removes and returns item from the front of the Queue.
//
// If the Queue is empty, panic occurs.
func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("can't dequeue, queue is empty")
	}
	q.mu.Lock()
	deq := q.items[0]
	q.items = q.items[1:]
	q.mu.Unlock()
	return deq
}

// DequeueAll removes and returns all items from the Queue.
func (q *Queue[T]) DequeueAll() []T {
	all := make([]T, 0)
	q.mu.Lock()
	for !q.Empty() {
		all = append(all, q.Dequeue())
	}
	q.mu.Unlock()
	return all
}

// Copy returns a shallow copy of the Queue.
func (q *Queue[T]) Copy() *Queue[T] {
	copy := New[T]()
	q.mu.Lock()
	for _, item := range q.items {
		copy.items = append(copy.items, item)
	}
	q.mu.Unlock()
	return copy
}

// Clear empties the Queue.
func (q *Queue[T]) Clear() {
	q.mu.Lock()
	q.items = make([]T, 0)
	q.mu.Unlock()
}

// Sprint formats items from the Queue using their default formats and returns the resulting string.
func (q *Queue[T]) Sprint() string {
	var sprint string
	q.mu.Lock()
	for _, item := range q.items {
		sprint += fmt.Sprintf("%v ", item)
	}
	q.mu.Unlock()
	return sprint
}

// Sprintf formats items from the Queue using formatter function and returns the resulting string.
func (q *Queue[T]) Sprintf(formatter func(item T) string) string {
	var sprint string
	q.mu.Lock()
	for _, item := range q.items {
		sprint += fmt.Sprintf("%s ", formatter(item))
	}
	q.mu.Unlock()
	return sprint
}

// Print formats items from the Queue using their default formats and writes resulting string to standart output.
func (q *Queue[T]) Print() {
	fmt.Printf("%s\n", q.Sprint())
}

// Printf formats items from the Queue using formatter function and writes resulting string to standart output.
func (q *Queue[T]) Printf(formatter func(item T) string) {
	fmt.Printf("%s\n", q.Sprintf(formatter))
}
