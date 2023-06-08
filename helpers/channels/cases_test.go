package channels

import (
	"context"
	"testing"

	. "github.com/vexgratia/collection-go/helpers/test"
)

type TestCase[T any] struct {
	name  string
	chans []chan T
	want  int
}

var intCases = []TestCase[int]{
	{"empty", make([]chan int, 0), 0},
	{"default", make([]chan int, 123), 123},
}
var stringCases = []TestCase[string]{
	{"empty", make([]chan string, 0), 0},
	{"default", make([]chan string, 45), 45},
}
var float64Cases = []TestCase[float64]{
	{"empty", make([]chan float64, 0), 0},
	{"default", make([]chan float64, 78), 78},
}

func TestMakeCases(t *testing.T) {
	t.Run("integer",
		func(t *testing.T) {
			for _, test := range intCases {
				cases := MakeCases(test.chans...)
				AssertEqual(t, len(cases), test.want)
			}
		})
	t.Run("string",
		func(t *testing.T) {
			for _, test := range stringCases {
				cases := MakeCases(test.chans...)
				AssertEqual(t, len(cases), test.want)
			}
		})
	t.Run("float64",
		func(t *testing.T) {
			for _, test := range float64Cases {
				cases := MakeCases(test.chans...)
				AssertEqual(t, len(cases), test.want)
			}
		})
}
func TestAddContext(t *testing.T) {
	ctx := context.Background()
	t.Run("integer",
		func(t *testing.T) {
			for _, test := range intCases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddContext(cases, ctx)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
	t.Run("string",
		func(t *testing.T) {
			for _, test := range stringCases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddContext(cases, ctx)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
	t.Run("float64",
		func(t *testing.T) {
			for _, test := range float64Cases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddContext(cases, ctx)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
}
func TestAddDefault(t *testing.T) {
	t.Run("integer",
		func(t *testing.T) {
			for _, test := range intCases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddDefault(cases)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
	t.Run("string",
		func(t *testing.T) {
			for _, test := range stringCases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddDefault(cases)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
	t.Run("float64",
		func(t *testing.T) {
			for _, test := range float64Cases {
				cases := MakeCases(test.chans...)
				ctxIndex := AddDefault(cases)
				AssertEqual(t, ctxIndex, test.want)
			}
		})
}
