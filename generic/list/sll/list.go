package list

// This file contains the implementation of generic SLL and its basic methods.

import (
	"sync"

	node "github.com/vexgratia/collection-go/generic/node/singly"
)

// A List is a generic implementation of SLL.
//
// List is based on singly linked Nodes of type T.
//
// List is thread-safe.
type List[T any] struct {
	mu     *sync.Mutex
	head   *node.Node[T]
	tail   *node.Node[T]
	length uint64
}

// New creates an empty List of type T.
func New[T any]() *List[T] {
	return &List[T]{
		mu: &sync.Mutex{},
	}
}

// Peek returns value of the head node of the list.
//
// If the List is empty, panic occurs.
func (l *List[T]) Peek() T {
	if l.Empty() {
		panic("can't peek, list is empty")
	}
	return l.head.Value
}

// Push inserts values to the List.
func (l *List[T]) Push(values ...T) {
	l.mu.Lock()
	for _, val := range values {
		node := node.New(val)
		if l.Len() == 0 {
			l.tail = node
		} else {
			node.Next = l.head
		}
		l.head = node
		l.length++
	}
	l.mu.Unlock()
}

// Trim removes head node from the List.
//
// If the List is empty, panic occurs.
func (l *List[T]) Trim() {
	if l.Empty() {
		panic("can't trim, list is empty")
	}
	l.mu.Lock()
	l.head = l.head.Next
	l.length--
	l.mu.Unlock()
}

// Collect returns values of all nodes from the List.
func (l *List[T]) Collect() []T {
	all := make([]T, 0)
	l.mu.Lock()
	current := l.head
	for current != nil {
		all = append(all, current.Value)
		current = current.Next
	}
	l.mu.Unlock()
	return all
}

// Copy returns a shallow copy of the List.
func (l *List[T]) Copy() *List[T] {
	copy := New[T]()
	all := l.Collect()
	for i := len(all) - 1; i >= 0; i-- {
		copy.Push(all[i])
	}
	return copy
}

// Clear empties the List.
func (l *List[T]) Clear() {
	l.mu.Lock()
	l.head = nil
	l.tail = nil
	l.length = 0
	l.mu.Unlock()
}
