package list

// This file contains the implementation of Printable interface methods for generic SCL.

import "fmt"

// Print formats values from the List using their default formats and writes resulting string to standart output.
func (l *List[T]) Print() {
	fmt.Printf("%s\n", l.Sprint())
}

// Printf formats values from the List using formatter function and writes resulting string to standart output.
func (l *List[T]) Printf(formatter func(value T) string) {
	fmt.Printf("%s\n", l.Sprintf(formatter))
}
