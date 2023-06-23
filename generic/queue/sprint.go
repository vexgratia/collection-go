package queue

// This file contains the implementation of Sprintable interface methods for generic Queue.

import "fmt"

// Sprint formats items from the Queue using their default formats and returns the resulting string.
func (q *Queue[T]) Sprint() string {
	var sprint string
	all := q.Collect()
	for _, item := range all {
		sprint += fmt.Sprintf("%v ", item)
	}
	return sprint
}

// Sprintf formats items from the Queue using formatter function and returns the resulting string.
func (q *Queue[T]) Sprintf(formatter func(item T) string) string {
	var sprint string
	all := q.Collect()
	for _, item := range all {
		sprint += fmt.Sprintf("%s ", formatter(item))
	}
	return sprint
}
