package queue

import "fmt"

type Queue[T any] struct {
	Data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
func (q *Queue[T]) Len() int {
	return len(q.Data)
}
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
func (q *Queue[T]) Peek() T {
	return q.Data[0]
}
func (q *Queue[T]) Enqueue(data T) {
	q.Data = append(q.Data, data)
}
func (q *Queue[T]) Dequeue() T {
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
func (q *Queue[T]) Print() {
	fmt.Printf("Queue: %v\n", q.Data)
}
func (q *Queue[T]) Copy() *Queue[T] {
	res := NewQueue[T]()
	copy(q.Data, res.Data)
	return res
}
func (q *Queue[T]) Sanitize() {
	q.Data = make([]T, 0)
}
