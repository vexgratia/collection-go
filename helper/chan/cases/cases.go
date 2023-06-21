package cases

import (
	"context"
	"reflect"
)

func MakeCases[T any](chans ...chan T) []reflect.SelectCase {
	cases := []reflect.SelectCase{}
	for _, ch := range chans {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)})
	}
	return cases
}
func AddContext(cases []reflect.SelectCase, ctx context.Context) ([]reflect.SelectCase, int) {
	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ctx.Done())})
	return cases, len(cases) - 1
}
func AddDefault(cases []reflect.SelectCase) ([]reflect.SelectCase, int) {
	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectDefault})
	return cases, len(cases) - 1
}
