package assert

import (
	"reflect"
	"testing"
)

func True(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func False(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
func Equal[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func NotEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("didn't want %v", got)
	}
}
func Panic(t *testing.T, panicFunc func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function didn't panic")
		}
	}()
	panicFunc()
}
func NotPanic(t *testing.T, panicFunc func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("function panics")
		}
	}()
	panicFunc()
}
