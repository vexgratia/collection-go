package node

// This file contains the implementation of generic singly linked Node and its basic methods.

// A Node represents a generic node in a singly linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// New creates a Node using value.
func New[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}
