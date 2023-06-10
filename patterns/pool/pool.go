package pool

import (
	"context"
	"sync/atomic"

	. "github.com/vexgratia/collection-go/patterns/task"
)

type PoolStatus int

const (
	POOL_STATUS_NEW PoolStatus = iota
	POOL_STATUS_READY
	POOL_STATUS_IN_PROCESS
	POOL_STATUS_DONE
	POOL_STATUS_RECOVERED
	POOL_STATUS_TERMINATED
	POOL_STATUS_TERMINATED_TIMEOUT
)

type Pool[I, O any] struct {
	ID      int
	Name    string
	State   *PoolState
	CTX     context.Context
	Workers []*Worker[I, O]
	Inputs  []chan *Task[I, O]
}
type PoolState struct {
	Status    PoolStatus
	Processed atomic.Uint64
}

func NewPool[I, O any](id int, name string) *Pool[I, O] {
	return &Pool[I, O]{
		ID:   id,
		Name: name,
		State: &PoolState{
			Status: POOL_STATUS_NEW,
		},
	}
}
