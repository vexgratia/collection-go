package list

type List[T any] interface {
	Peek() T
	Push(values ...T)
	Trim()
}

type Circular interface {
	Scroll()
}
