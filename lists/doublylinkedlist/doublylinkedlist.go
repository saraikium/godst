package doublylinkedlist

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
