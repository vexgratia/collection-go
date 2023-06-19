package sll

import (
	"fmt"
	"sync"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}
func (l *List[T]) IsHead(node *Node[T]) bool {
	return l.Head == node
}
func (l *List[T]) IsTail(node *Node[T]) bool {
	return l.Tail == node
}
func (l *List[T]) InMiddle(node *Node[T]) bool {
	return !l.IsHead(node) && !l.IsTail(node)
}

type List[T any] struct {
	Mu     *sync.Mutex
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func New[T any]() *List[T] {
	return &List[T]{
		Mu: &sync.Mutex{},
	}
}
func (l *List[T]) Len() int {
	return l.Length
}

func (l *List[T]) Push(value T) {
	node := NewNode(value)
	l.Mu.Lock()
	if l.Len() == 0 {
		l.Tail = node
	} else {
		node.Next = l.Head
	}
	l.Head = node
	l.Length++
	l.Mu.Unlock()
}
func (l *List[T]) Trim() {
	l.Mu.Lock()
	l.Head = l.Head.Next
	l.Length--
	l.Mu.Unlock()
}

func (l *List[T]) Clear() {
	l.Mu.Lock()
	l.Head = nil
	l.Tail = nil
	l.Length = 0
	l.Mu.Unlock()
}

func (l *List[T]) Print() {
	fmt.Printf("List: ")
	l.Mu.Lock()
	node := l.Head
	for node != nil {
		fmt.Printf("%v", node.Value)
		node = node.Next
		if node != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Printf("\n")
	l.Mu.Unlock()
}
