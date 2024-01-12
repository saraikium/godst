package doublylinkedlist

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	size int
	tail *Node[T]
}

func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}

	return list
}

func (list *List[T]) Add(values ...T) {

	for _, value := range values {
		node := &Node[T]{value: value, prev: list.tail}
		if list.size == 0 {
			list.head = node
			list.tail = node
		} else {
			list.tail.next = node
			list.tail = node
		}
		list.size += 1
	}

}

func (list *List[T]) InsertAt(index int, value T) (bool, error) {

	if index > list.size {
		return false, fmt.Errorf("Invalid index. Out of bounds.")
	}

	if index < 0 {
		return false, fmt.Errorf("Index can't be negative.")
	}

	if index == list.size {
		list.Add(value)
		return true, nil
	}

	var runner *Node[T] = list.head

	for i := 0; i < index-1 && runner.next != nil; i++ {
		runner = runner.next
	}

	node := &Node[T]{value: value, next: runner.next, prev: runner}

	runner.next = node
	node.next.prev = node
	list.size += 1

	runner = nil

	return true, nil

}

func (list *List[T]) RemoveAt(index int) (T, error) {

	if index >= list.size || index < 0 {
		return *new(T), fmt.Errorf("Invalid index. Either less than 0 or out of bound.")
	}
	runner := list.head

	for i := 0; i < index; i, runner = i+1, runner.next {
	}

	if runner == list.head {
		list.head = runner.next
	}

	if runner == list.tail {
		list.tail = runner.prev
	}

	if runner.next != nil {
		runner.next.prev = runner.prev
	}

	if runner.prev != nil {
		runner.prev.next = runner.next
	}

	v := runner.value
	runner = nil

	list.size -= 1

	return v, nil
}

func (list *List[T]) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list *List[T]) Contains(element T, checker func(T) bool) bool {
	for element := list.head; element != nil; element = element.next {
		if checker(element.value) {
			return true
		}
	}
	return false
}

func (list *List[T]) Empty() bool { return list.size == 0 }

func (list *List[T]) Size() int { return list.size }

func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := make([]string, 0, list.size)
	for element := list.head; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}
