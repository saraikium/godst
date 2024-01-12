package stacks

type Stack[T any] interface {
	Push(value T) error
	Pop() (T, error)
}
