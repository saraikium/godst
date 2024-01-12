package dynamicarray

import (
	"fmt"
)

type DynamicArray[T any] struct {
	elements []T
}

func New[T any](values ...T) *DynamicArray[T] {

	arr := &DynamicArray[T]{}
	arr.elements = make([]T, 0, 0)

	if len(values) > 0 {
		arr.Add(values...)
	}

	return arr
}

func (arr *DynamicArray[T]) Size() int {
	return len(arr.elements)
}

func (arr *DynamicArray[T]) Empty() bool {
	return len(arr.elements) == 0
}

func (arr *DynamicArray[T]) Get(index int) (*T, error) {
	if index >= len(arr.elements) {
		return nil, fmt.Errorf("Index %v out of bounds. len is %v", index, len(arr.elements))
	}
	return &arr.elements[index], nil
}

func (arr *DynamicArray[T]) Set(index int, value T) (bool, error) {

	if index >= len(arr.elements) {
		return false, fmt.Errorf("Index: %v out of bounds. Length is %v", index, len(arr.elements))
	}

	arr.elements[index] = value
	return true, nil
}

func (arr *DynamicArray[T]) Add(values ...T) {
	arr.elements = append(arr.elements, values...)
}

func (arr *DynamicArray[T]) RemoveAt(index int) (*T, error) {

	value := &arr.elements[index]
	if index >= len(arr.elements) {
		return nil, fmt.Errorf("Index %v out of bounds. len is %v", index, len(arr.elements))
	}
	clear(arr.elements[index : index+1])
	copy(arr.elements[index:], arr.elements[index+1:arr.Size()])
	return value, nil

}

func (arr *DynamicArray[T]) IndexOf(target T, comparator func(T, T) bool) (int, error) {

	for index, element := range arr.elements {
		if comparator(target, element) {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Not found.")
}
