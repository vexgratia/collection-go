package generator

import (
	"context"
	"sync/atomic"

	. "github.com/vexgratia/collection-go/patterns/task"
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

type Generator[I, O any] struct {
	ID       int
	Name     string
	State    *GeneratorState
	CTX      context.Context
	Threads  []*Thread[I, O]
	Generate func(current uint64, example I) (*Task[I, O], error)
	Example  I
}
type GeneratorState struct {
	Status  GeneratorStatus
	Current atomic.Uint64
	Goal    atomic.Uint64
}

func NewGenerator[I, O any](id int, name string) *Generator[I, O] {
	return &Generator[I, O]{
		ID:   id,
		Name: name,
		State: &GeneratorState{
			Status: GENERATOR_STATUS_NEW,
		},
	}
}
func (g *Generator[I, O]) Size() int {
	return len(g.Threads)
}
