package arraystack

import (
	"fmt"
	"github.com/saraikium/godst/stacks"
)

type Stack[T any] struct {
	data  []T
	limit int
}

var _ stacks.Stack[int] = (*Stack[int])(nil)

func New[T any](limit int) *Stack[T] {
	s := &Stack[T]{}

	s.limit = limit
	s.data = make([]T, 0, limit)

	return s
}

func (s *Stack[T]) Empty() bool {
	return s.Size() == 0

}
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Search(cb func(T) bool) bool {
	for _, item := range s.data {
		if cb(item) {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Push(v T) error {
	if s.Size() >= s.limit {
		return fmt.Errorf("Can't grow the stack anymore. Limit is %v elements.", s.limit)
	}
	s.data = append(s.data, v)
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		return *new(T), fmt.Errorf("Empty Stack.")
	}
	element := s.data[s.Size()-1]

	clear(s.data[s.Size()-1:])
	s.data = s.data[:s.Size()-1]

	return element, nil
}

func (s *Stack[T]) Peek() T {
	v := s.data[s.Size()-1]
	return v
}
