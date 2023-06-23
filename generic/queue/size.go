package queue

// This file contains the implementation of Sized interface methods for generic Queue.

// Len returns current length of the Queue.
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// Empty returns true if the Queue is empty, and false if not.
func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}
