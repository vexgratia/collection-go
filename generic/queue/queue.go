package queue

// This file contains the implementation of generic Queue and its basic methods.

import (
	"sync"
)

// A Queue is a generic linear data structure.
//
// Queue is based on slice of type T.
//
// Queue is thread-safe.
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

// Peek returns the last item that's been enqueued.
//
// If the Queue is empty, panic occurs.
func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("can't peek, queue is empty")
	}
	return q.items[0]
}

// Enqueue inserts values in the end of the Queue.
func (q *Queue[T]) Enqueue(values ...T) {
	q.mu.Lock()
	for _, val := range values {
		q.items = append(q.items, val)
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

// Collect returns all items from the Queue.
func (q *Queue[T]) Collect() []T {
	return q.items
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
