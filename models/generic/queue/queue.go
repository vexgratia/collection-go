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

// Peek returns the last item that's been Enqueued.
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

// Enqueue inserts value in the end of the Queue.
func (q *Queue[T]) Enqueue(value T) {
	q.mu.Lock()
	q.items = append(q.items, value)
	q.mu.Unlock()
}

// Enqueue inserts multiple values in the end of the Queue.
func (q *Queue[T]) EnqueueAll(values ...T) {
	for _, val := range values {
		q.Enqueue(val)
	}
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

// DequeueAll removes and returns all values from the Queue.
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

// Print writes all items from the queue to standard output.
func (q *Queue[T]) Print() {
	for _, item := range q.items {
		fmt.Printf("%v ", item)
	}
	fmt.Printf("\n")
}

// Printf calls formatter function on each item of the Queue and writes results to standard output.
func (q *Queue[T]) Printf(formatter func(T) string) {
	for _, item := range q.items {
		fmt.Printf("%s ", formatter(item))
	}
	fmt.Printf("\n")
}
