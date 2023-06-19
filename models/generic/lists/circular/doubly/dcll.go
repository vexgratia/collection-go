package dcll

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
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
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func New[T any]() *List[T] {
	return &List[T]{}
}
func (l *List[T]) Len() int {
	return l.Length
}

func (l *List[T]) Push(value T) {
	node := NewNode(value)
	if l.Len() == 0 {
		l.Tail = node
		node.Next = node
		node.Prev = node
	} else {
		node.Next = l.Head
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Head.Prev = node
	}
	l.Head = node
	l.Length++
}
func (l *List[T]) Trim() {
	l.Head = l.Head.Next
	l.Head.Prev = l.Tail
	l.Tail.Next = l.Head
	l.Length--
}
func (l *List[T]) ScrollNext() {
	l.Tail, l.Head = l.Head, l.Head.Next
}
func (l *List[T]) ScrollPrev() {
	l.Tail, l.Head = l.Tail.Prev, l.Tail
}
func (l *List[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Length = 0
}

func (l *List[T]) Print() {
	fmt.Printf("List: ")
	if l.Len() == 0 {
		fmt.Printf("\n")
		return
	}
	fmt.Printf(" ... <-> ")
	head := l.Head
	node := head
	for {
		fmt.Printf("%v", node.Value)
		node = node.Next
		if node == head {
			fmt.Printf(" <-> ...\n")
			return
		} else {
			fmt.Printf(" <-> ")
		}
	}
}
