package queues

import (
	"testing"
)

func TestEnqueue(t *testing.T) {

	q := New[int]()

	_, err := q.Dequeue()

	if err == nil {
		t.Errorf("Empty queue should throw an error on Dequeue.")
	}

	q.Enqueue(5)
	q.Enqueue(6)

	if q.Size() != 2 {
		t.Errorf("Size should be 2. Got=%v", q.Size())
	}

	v, err := q.Dequeue()

	if q.Size() != 1 {
		t.Errorf("Size should be 1. got %v", q.Size())
	}

	if err != nil {
		t.Errorf("Value should be 5, got=%v", v)
	}

	v, err = q.Dequeue()

	if q.Size() != 0 {
		t.Errorf("Size should be 1. got %v", q.Size())
	}

	if err != nil {
		t.Errorf("Value should be 6, got=%v", v)
	}

	_, err = q.Dequeue()

	if q.Size() != 0 {
		t.Errorf("Size should be 1. got %v", q.Size())
	}

	if err == nil {
		t.Errorf("Empty queue should throw an error on Dequeue.")
	}

}
