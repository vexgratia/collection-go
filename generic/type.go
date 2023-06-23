package generic

type Sized interface {
	Len() int
	Empty() bool
}
type Sprintable interface {
	Sprint() string
	Sprintf() string
}
type Printable interface {
	Print()
	Printf()
}
