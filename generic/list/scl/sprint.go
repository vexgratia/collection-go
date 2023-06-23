package list

// This file contains the implementation of Sprintable interface methods for generic SCL.

import "fmt"

// Sprint formats values from the List using their default formats and returns the resulting string.
func (l *List[T]) Sprint() string {
	var sprint string
	if l.Empty() {
		return sprint
	}
	l.mu.Lock()
	sprint += fmt.Sprintf("... -> %v", l.head.Value)
	current := l.head.Next
	for current != l.head {
		sprint += fmt.Sprintf(" -> %v", current.Value)
		current = current.Next
	}
	sprint += fmt.Sprintf(" -> ...")
	l.mu.Unlock()
	return sprint
}

// Sprintf formats values from the List using formatter function and returns the resulting string.
func (l *List[T]) Sprintf(format func(value T) string) string {
	var sprint string
	if l.Empty() {
		return sprint
	}
	l.mu.Lock()
	sprint += fmt.Sprintf("... -> %s", format(l.head.Value))
	current := l.head.Next
	for current != l.head {
		sprint += fmt.Sprintf(" -> %s", format(current.Value))
		current = current.Next
	}
	sprint += fmt.Sprintf(" -> ...")
	l.mu.Unlock()
	return sprint
}
