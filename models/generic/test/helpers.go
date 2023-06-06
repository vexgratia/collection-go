package test

import (
	"reflect"
	"testing"
)

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
func AssertEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("didn't want %v", got)
	}
}
func AssertPanic(t *testing.T, panicFunc func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function didn't panic")
		}
	}()
	panicFunc()
}
func AssertNotPanic(t *testing.T, panicFunc func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("function panics")
		}
	}()
	panicFunc()
}
