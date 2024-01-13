package buffers

import (
	"sync"
)

type RingBuffer[T any] struct {
	data  []T
	head  int
	mutex sync.Mutex
	size  int
	tail  int
}

func New[T any](size int) *RingBuffer[T] {
	r := &RingBuffer[T]{
		data: make([]T, size),
		size: size,
	}

	return r
}
