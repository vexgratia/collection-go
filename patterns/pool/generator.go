package pool

import (
	"context"
	"sync/atomic"
)

type GeneratorStatus int

const (
	GENERATOR_STATUS_NEW GeneratorStatus = iota
	GENERATOR_STATUS_READY
	GENERATOR_STATUS_IN_PROCESS
	GENERATOR_STATUS_DONE
	GENERATOR_STATUS_RECOVERED
	GENERATOR_STATUS_TERMINATED
	GENERATOR_STATUS_TERMINATED_TIMEOUT
)

type Generator[T any] struct {
	ID       int
	Name     string
	State    *GeneratorState
	CTX      context.Context
	Threads  []*Thread[T]
	Generate func(current uint64, example T) (*Task[T], error)
	Example  T
}
type GeneratorState struct {
	Status  GeneratorStatus
	Current atomic.Uint64
	Goal    atomic.Uint64
}

func NewGenerator[T any](id int, name string) *Generator[T] {
	return &Generator[T]{
		ID:   id,
		Name: name,
		State: &GeneratorState{
			Status: GENERATOR_STATUS_NEW,
		},
	}
}
func (g *Generator[T]) Size() int {
	return len(g.Threads)
}
