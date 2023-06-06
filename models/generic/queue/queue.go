package queue

import "fmt"

type Queue[T any] struct {
	Data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
func (q *Queue[T]) len() int {
	return len(q.Data)
}
func (q *Queue[T]) isEmpty() bool {
	return q.len() == 0
}
func (q *Queue[T]) peek() T {
	return q.Data[0]
}
func (q *Queue[T]) enqueue(data T) {
	q.Data = append(q.Data, data)
}
func (q *Queue[T]) dequeue() T {
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
func (q *Queue[T]) print() {
	fmt.Printf("Queue: %v\n", q.Data)
}
func (q *Queue[T]) copy() *Queue[T] {
	res := NewQueue[T]()
	for _, data := range q.Data {
		res.Data = append(res.Data, data)
	}
	return res
}
func (q *Queue[T]) sanitize() {
	q.Data = make([]T, 0)
}
