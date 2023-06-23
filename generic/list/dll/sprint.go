package list

// This file contains the implementation of Sprintable interface methods for generic DLL.

import "fmt"

// Sprint formats values from the List using their default formats and returns the resulting string.
func (l *List[T]) Sprint() string {
	var sprint string
	if l.Empty() {
		return sprint
	}
	l.mu.Lock()
	current := l.head
	for current != nil {
		sprint += fmt.Sprintf(" <-> %v", current.Value)
		current = current.Next
	}
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
	current := l.head
	for current != nil {
		sprint += fmt.Sprintf(" <-> %s", format(current.Value))
		current = current.Next
	}
	l.mu.Unlock()
	return sprint
}
