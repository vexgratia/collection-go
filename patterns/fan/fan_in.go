package fan

import (
	"context"
	"reflect"
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
