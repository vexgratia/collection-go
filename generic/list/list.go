package list

type List[T any] interface {
	Peek() T
	Push(values ...T)
	Trim()
	Collect()
	Copy()
	Clear()
}
