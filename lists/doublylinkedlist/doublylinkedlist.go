package doublylinkedlist

import (
	"github.com/saraikium/godst/lists"
)

// Type assertion. Checks if DoublyLinkedList implements all the methods of List
var _ lists.List[any] = (*DoublyLinkedList[any])(nil)

type Node[T any] struct {
	value    T
	next     *Node[T]
	previous *Node[T]
}

type DoublyLinkedList[T any] struct {
	head *Node[T]
	size int
	tail *Node[T]
}

func newNode[T any](v T) *Node[T] {
	node := &Node[T]{value: v, next: nil, previous: nil}
	return node
}

func New()
