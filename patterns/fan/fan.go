package fan

import (
	"context"
	"reflect"

	. "github.com/vexgratia/collection-go/helpers/channels"
)

func FanIn[T any](ctx context.Context, output chan T, inputs ...chan T) {
	cases := MakeCases(inputs...)
	cases, ctxIndex := AddContext(cases, ctx)
	cases, defaultIndex := AddDefault(cases)
	for {
		index, data, _ := reflect.Select(cases)
		switch {
		case index < ctxIndex:
			output <- data.Interface().(T)
		case index == ctxIndex:
			break
		case index == defaultIndex:
			continue
		}
	}
}
func FanOut[T any](ctx context.Context, input chan T, outputs ...chan T) {
	for {
		select {
		case data := <-input:
			for _, output := range outputs {
				output <- data
			}
		case <-ctx.Done():
			break
		default:
			continue
		}
	}
}
