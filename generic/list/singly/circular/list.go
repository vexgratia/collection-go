package list

import (
	"fmt"
	"sync"

	node "github.com/vexgratia/collection-go/generic/list/singly"
)

// A List is a generic linear data structure.
//
// List is based on Nodes of type T.
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

// Len returns current length of the List.
func (l *List[T]) Len() int {
	return int(l.length)
}

// Empty returns true if the List is empty, and false if not.
func (l *List[T]) Empty() bool {
	return l.Len() == 0
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

// PeekAll returns values of all nodes from the List.
func (l *List[T]) PeekAll() []T {
	all := make([]T, 0)
	l.mu.Lock()
	current := l.head
	for {
		all = append(all, current.Value)
		current := current.Next
		if current == l.head {
			break
		}
	}
	l.mu.Unlock()
	return all
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
		l.tail.Next = l.head
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
	l.tail.Next = l.head.Next
	l.head = l.head.Next
	l.length--
	l.mu.Unlock()
}

// Clear empties the List.
func (l *List[T]) Clear() {
	l.mu.Lock()
	l.head = nil
	l.tail = nil
	l.length = 0
	l.mu.Unlock()
}

// Sprint formats values from the List using their default formats and returns the resulting string.
func (l *List[T]) Sprint() string {
	var sprint string
	if l.Empty() {
		return sprint
	}
	all := l.PeekAll()
	sprint += "... -> "
	for _, val := range all {
		sprint += fmt.Sprintf("%v -> ", val)
	}
	sprint += "..."
	return sprint
}

// Sprintf formats values from the List using formatter function and returns the resulting string.
func (l *List[T]) Sprintf(formatter func(value T) string) string {
	var sprint string
	if l.Empty() {
		return sprint
	}
	all := l.PeekAll()
	sprint += "... -> "
	for _, val := range all {
		sprint += fmt.Sprintf("%s -> ", formatter(val))
	}
	sprint += "..."
	return sprint
}

// Print formats values from the List using their default formats and writes resulting string to standart output.
func (l *List[T]) Print() {
	fmt.Printf("%s\n", l.Sprint())
}

// Printf formats values from the List using formatter function and writes resulting string to standart output.
func (l *List[T]) Printf(formatter func(value T) string) {
	fmt.Printf("%s\n", l.Sprintf(formatter))
}
