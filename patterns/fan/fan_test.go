package fan

import (
	"context"
	"testing"
)

func TestFanIn(t *testing.T) {
	type TestCase struct {
		name   string
		output chan int
		inputs []chan int
	}
	cases := []TestCase{
		{"empty", make(chan int), []chan int{}},
		{"default", make(chan int, 10), []chan int{make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)}},
	}
	for _, test := range cases {
		ctx, done := context.WithCancel(context.Background())
		dic := make(map[int]bool)
		go FanIn(ctx, test.output, test.inputs...)
		for id, inp := range test.inputs {
			inp <- id
			dic[id] = true
		}
		for range test.inputs {
			got := <-test.output
			delete(dic, got)
		}
		if len(dic) > 0 {
			t.Errorf("missing %d values", len(dic))
		}
		done()
	}
}
func TestFanOut(t *testing.T) {
	type TestCase struct {
		name    string
		input   chan int
		outputs []chan int
		want    int
	}
	cases := []TestCase{
		{"empty", make(chan int), []chan int{}, 123},
		{"default", make(chan int, 10), []chan int{make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)}, 789},
	}
	for _, test := range cases {
		ctx, done := context.WithCancel(context.Background())
		go FanOut(ctx, test.input, test.outputs...)
		test.input <- test.want
		for _, output := range test.outputs {
			got := <-output
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		}
		done()
	}
}
