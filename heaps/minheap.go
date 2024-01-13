package heaps

import (
	"cmp"
	"fmt"
)

type MinHeap[T cmp.Ordered] struct {
	size  int
	items []T
}

func New[T cmp.Ordered]() *MinHeap[T] {
	h := &MinHeap[T]{}
	h.size = 0
	h.items = make([]T, 0, 2)
	return h
}

func (h *MinHeap[T]) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}
func (h *MinHeap[T]) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}
func (h *MinHeap[T]) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}
func (h *MinHeap[T]) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.size
}
func (h *MinHeap[T]) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.size
}
func (h *MinHeap[T]) hasParent(index int) bool {
	return h.getParentIndex(index) < h.size
}

func (h *MinHeap[T]) leftChild(index int) T {
	return h.items[h.getLeftChildIndex(index)]
}
func (h *MinHeap[T]) rightChild(index int) T {
	return h.items[h.getRightChildIndex(index)]
}
func (h *MinHeap[T]) parent(index int) T {
	return h.items[h.getParentIndex(index)]
}

func (h *MinHeap[T]) swap(firstIndex, secondIndex int) {
	h.items[firstIndex], h.items[secondIndex] = h.items[secondIndex], h.items[firstIndex]
}

func (h *MinHeap[T]) reduceCapacity() {
	targetCapacity := int(float64(cap(h.items)) * 0.66)

	if h.size < targetCapacity {
		items := make([]T, h.size, targetCapacity)
		copy(items, h.items[:h.size])
		h.items = items
	}
}

func (h *MinHeap[T]) Peek() (T, error) {
	if h.size == 0 {
		return *new(T), fmt.Errorf("Empty Heap.")
	}
	return h.items[0], nil
}

func (h *MinHeap[T]) Poll() (T, error) {

	if h.size == 0 {
		return *new(T), fmt.Errorf("Empty Heap.")
	}

	item := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--

	h.heapifyDown()
	h.reduceCapacity()
	return item, nil
}

func (h *MinHeap[T]) Add(item T) {
	h.items = append(h.items, item)
	h.size++
	h.heapifyUp()
}

func (h *MinHeap[T]) heapifyUp() {
	index := h.size - 1

	for h.hasParent(index) && h.parent(index) > h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}

}

func (h *MinHeap[T]) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)

		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = h.getRightChildIndex(index)
		}

		if h.items[index] <= h.items[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
		}

		index = smallerChildIndex
	}

}
