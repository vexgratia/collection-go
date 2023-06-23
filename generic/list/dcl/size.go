package list

// This file contains the implementation of Sized interface methods for generic DCL.

// Len returns current length of the List.
func (l *List[T]) Len() int {
	return int(l.length)
}

// Empty returns true if the List is empty, and false if not.
func (l *List[T]) Empty() bool {
	return l.Len() == 0
}
