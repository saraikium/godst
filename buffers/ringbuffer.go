package buffers

import (
	"fmt"
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

func (rb *RingBuffer[T]) Enqueue(v T) error {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	if (rb.tail+1)%rb.size == rb.head {
		return fmt.Errorf("RingBuffer is full.")
	}

	rb.data[rb.tail] = v
	rb.tail = (rb.tail + 1) % rb.size
	return nil
}

func (rb *RingBuffer[T]) Deque() (T, error) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	var zeroVal T
	if rb.tail == rb.head {
		return zeroVal, fmt.Errorf("Empty Buffer")
	}
	data := rb.data[rb.head]
	rb.head = (rb.head + 1) % rb.size

	return data, nil
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.head == r.tail
}

func (r *RingBuffer[T]) IsFull() bool {
	return (r.tail+1)%r.size == r.head
}
