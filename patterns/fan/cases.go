package fan

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
func AddContext(cases []reflect.SelectCase, ctx context.Context) int {
	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ctx.Done())})
	return len(cases) - 1
}
func AddDefault(cases []reflect.SelectCase) int {
	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectDefault})
	return len(cases) - 1
}
