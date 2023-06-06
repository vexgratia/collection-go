package test

import (
	"testing"
)

func TestAssertTrue(t *testing.T) {
	type TestCase struct {
		input bool
	}
	cases := []TestCase{
		{true},
	}
	for _, test := range cases {
		AssertTrue(t, test.input)
	}
}
func TestAssertFalse(t *testing.T) {
	type TestCase struct {
		input bool
	}
	cases := []TestCase{
		{false},
	}
	for _, test := range cases {
		AssertFalse(t, test.input)
	}
}
func TestAssertEqual(t *testing.T) {
	type TestCase[T any] struct {
		input T
		want  T
	}
	t.Run("integer", func(t *testing.T) {
		cases := []TestCase[int]{
			{0, 0},
			{1234, 1234},
			{-5678, -5678},
		}
		for _, test := range cases {
			AssertEqual(t, test.input, test.want)
		}
	})
	t.Run("string", func(t *testing.T) {
		cases := []TestCase[string]{
			{"", ""},
			{"abc123!@#", "abc123!@#"}}
		for _, test := range cases {
			AssertEqual(t, test.input, test.want)
		}
	})
	t.Run("float64", func(t *testing.T) {
		cases := []TestCase[float64]{
			{0.0, 0.0},
			{123.456, 123.456},
			{-456.789, -456.789},
		}
		for _, test := range cases {
			AssertEqual(t, test.input, test.want)
		}
	})
	t.Run("map", func(t *testing.T) {
		cases := []TestCase[map[string]int]{
			{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}},
		}
		for _, test := range cases {
			AssertEqual(t, test.input, test.want)
		}
	})
	t.Run("slice", func(t *testing.T) {
		cases := []TestCase[[]int]{
			{[]int{1, 2, 3}, []int{1, 2, 3}},
		}
		for _, test := range cases {
			AssertEqual(t, test.input, test.want)
		}
	})
}
func TestAssertNotEqual(t *testing.T) {
	type TestCase[T any] struct {
		input T
		want  T
	}
	t.Run("integer", func(t *testing.T) {
		cases := []TestCase[int]{
			{0, 1},
			{1234, -1234},
			{-5678, -5679},
		}
		for _, test := range cases {
			AssertNotEqual(t, test.input, test.want)
		}
	})
	t.Run("string", func(t *testing.T) {
		cases := []TestCase[string]{
			{"", "."},
			{"abc123!@#", "abc123!@# "}}
		for _, test := range cases {
			AssertNotEqual(t, test.input, test.want)
		}
	})
	t.Run("float64", func(t *testing.T) {
		cases := []TestCase[float64]{
			{0.0, 0.01},
			{123.456, 1234.56},
			{-456.789, 456.789},
		}
		for _, test := range cases {
			AssertNotEqual(t, test.input, test.want)
		}
	})
	t.Run("map", func(t *testing.T) {
		cases := []TestCase[map[string]int]{
			{map[string]int{"a": 1, "b": 2}, map[string]int{"A": 1, "b": 2}},
		}
		for _, test := range cases {
			AssertNotEqual(t, test.input, test.want)
		}
	})
	t.Run("slice", func(t *testing.T) {
		cases := []TestCase[[]int]{
			{[]int{1, 2, 3}, []int{-1, 2, 3}},
		}
		for _, test := range cases {
			AssertNotEqual(t, test.input, test.want)
		}
	})

}
func TestAssertPanic(t *testing.T) {
	type TestCase struct {
		input func()
	}
	cases := []TestCase{
		{func() {
			slice := []int{}
			slice[0] += 2
		}},
		{func() {
			slice := []int{1, 2, 3}
			slice[4] += 2
		}},
		{func() {
			slice := []int{1, 2, 3}
			index := -1
			slice[index] += 2
		}},
	}
	for _, test := range cases {
		AssertPanic(t, func() { test.input() })
	}
}
func TestAssertNotPanic(t *testing.T) {
	type TestCase struct {
		input func()
	}
	cases := []TestCase{
		{func() {
			slice := []int{0}
			slice[0] += 2
		}},
		{func() {
			slice := []int{1, 2, 3}
			slice[2] += 2
		}},
		{func() {
			slice := []int{1, 2, 3}
			index := 2
			slice[index] += 2
		}},
	}
	for _, test := range cases {
		AssertNotPanic(t, func() { test.input() })
	}
}
