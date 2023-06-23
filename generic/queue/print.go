package queue

// This file contains the implementation of Printable interface methods for generic Queue.

import "fmt"

// Print formats items from the Queue using their default formats and writes resulting string to standart output.
func (q *Queue[T]) Print() {
	fmt.Printf("%s\n", q.Sprint())
}

// Printf formats items from the Queue using formatter function and writes resulting string to standart output.
func (q *Queue[T]) Printf(formatter func(item T) string) {
	fmt.Printf("%s\n", q.Sprintf(formatter))
}
