package generic

type EqualFunc[T any] func(a, b T) bool

type LessFunc[T any] func(a, b T) bool

type MinFunc[T any] func(values ...T) T

type MaxFunc[T any] func(values ...T) T

func MakeMinFunc[T any](less LessFunc[T]) MinFunc[T] {
	return func(values ...T) T {
		if len(values) == 0 {
			panic("can't min, no values")
		}
		min := values[0]
		for _, val := range values {
			if less(val, min) {
				min = val
			}
		}
		return min
	}
}

func MakeMaxFunc[T any](less LessFunc[T]) MaxFunc[T] {
	return func(values ...T) T {
		if len(values) == 0 {
			panic("can't max, no values")
		}
		max := values[0]
		for _, val := range values {
			if less(max, val) {
				max = val
			}
		}
		return max
	}
}

func Occur[T any](equal EqualFunc[T], target T, values ...T) bool {
	for _, val := range values {
		if equal(val, target) {
			return true
		}
	}
	return false
}
