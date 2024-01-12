package stacks

type Stack[T any] interface {
	Empty() bool
	Search(cb func(T) bool) bool
	Size() int
	Pop() (T, error)
	Push(value T) error
	Peek() T
}
