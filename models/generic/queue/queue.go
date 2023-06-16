package queue

import (
	"fmt"
	"sync"
)

type Queue[T any] struct {
	Mu   *sync.Mutex
	Data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Mu: &sync.Mutex{},
	}
}
func (q *Queue[T]) Len() int {
	return len(q.Data)
}
func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}
func (q *Queue[T]) Peek() T {
	return q.Data[0]
}
func (q *Queue[T]) Enqueue(data T) {
	q.Mu.Lock()
	q.Data = append(q.Data, data)
	q.Mu.Unlock()
}
func (q *Queue[T]) Dequeue() T {
	q.Mu.Lock()
	res := q.Data[0]
	q.Data = q.Data[1:]
	q.Mu.Unlock()
	return res
}
func (q *Queue[T]) Copy() *Queue[T] {
	res := NewQueue[T]()
	q.Mu.Lock()
	for _, data := range q.Data {
		res.Data = append(res.Data, data)
	}
	q.Mu.Unlock()
	return res
}
func (q *Queue[T]) Clear() {
	q.Mu.Lock()
	q.Data = nil
	q.Mu.Unlock()
}
func (q *Queue[T]) Print() {
	fmt.Printf("Queue: %v\n", q.Data)
}
