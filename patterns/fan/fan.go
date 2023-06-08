package fan

import (
	"context"
	"reflect"

	. "github.com/vexgratia/collection-go/helpers/channels"
)

func FanIn[T any](ctx context.Context, output chan T, inputs ...chan T) {
	cases := MakeCases(inputs...)
	ctxIndex := AddContext(cases, ctx)
	defaultIndex := AddDefault(cases)
	for {
		switch index, data, _ := reflect.Select(cases); index {
		case ctxIndex:
			break
		case defaultIndex:
			continue
		default:
			output <- data.Interface().(T)
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
