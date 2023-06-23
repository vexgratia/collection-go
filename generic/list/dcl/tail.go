package list

// This file contains the implementation of generic DCL methods, starting from tail node.

// PeekTail returns value of the tail node of the list.
//
// If the List is empty, panic occurs.
func (l *List[T]) PeekTail() T {
	if l.Empty() {
		panic("can't peek, list is empty")
	}
	return l.tail.Value
}

// TrimTail removes tail node from the List.
//
// If the List is empty, panic occurs.
func (l *List[T]) TrimTail() {
	if l.Empty() {
		panic("can't trim, list is empty")
	}
	l.mu.Lock()
	l.tail = l.tail.Prev
	l.tail.Next = l.head
	l.length--
	l.mu.Unlock()
}
