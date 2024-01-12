package queues

import "fmt"

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	q := &Queue[T]{}
	q.data = make([]T, 0, 1)
	return q
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Peek() T {
	return q.data[q.Size()-1]
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Empty() {
		return *new(T), fmt.Errorf("Queue is Empty")
	}
	v := q.data[0]
	q.data = q.data[:q.Size()-1]
	return v, nil
}
