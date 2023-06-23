package list

import node "github.com/vexgratia/collection-go/generic/node/doubly"

// This file contains the implementation of generic DLL methods, starting from tail node.

// PeekTail returns value of the tail node of the list.
//
// If the List is empty, panic occurs.
func (l *List[T]) PeekTail() T {
	if l.Empty() {
		panic("can't peek, list is empty")
	}
	return l.tail.Value
}

// PushTail inserts values to the List.
func (l *List[T]) PushTail(values ...T) {
	l.mu.Lock()
	for _, val := range values {
		node := node.New(val)
		if l.Len() == 0 {
			node.Next = node
			node.Prev = node
			l.head = node
		} else {
			node.Prev = l.tail
			l.tail.Next = node
		}
		l.tail = node
		l.length++
	}
	l.mu.Unlock()
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
	l.tail.Next = nil
	l.length--
	l.mu.Unlock()
}
