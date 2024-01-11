package dynamicarray

import "fmt"

type DynamicArray[T any] struct {
	elments  []T
	len      int
	capacity int
}

func New[T any](values ...T) *DynamicArray[T] {

	arr := &DynamicArray[T]{}
	arr.len = 0
	arr.capacity = 0
	arr.elments = make([]T, 0, 0)

	if len(values) > 0 {
		arr.Add(values...)
	}

	return arr
}

func (arr *DynamicArray[T]) Size() int {
	return arr.len
}

func (arr *DynamicArray[T]) Empty() bool {
	return arr.len == 0
}

func (arr *DynamicArray[T]) Get(index int) (*T, error) {
	if index >= arr.len {
		return nil, fmt.Errorf("Index %v out of bounds. len is %v", index, arr.len)
	}
	return &arr.elments[index], nil
}

func (arr *DynamicArray[T]) Set(index int, value T) (bool, error) {

	if index >= arr.len {
		return false, fmt.Errorf("Index: %v out of bounds. Length is %v", index, arr.len)
	}

	arr.elments[index] = value
	return true, nil
}

func (arr *DynamicArray[T]) Add(values ...T) {
	toAdd := len(values)
	if arr.capacity <= arr.len+toAdd {
		arr.capacity += (arr.len + toAdd) * 2
	}
	arr.len += arr.len + toAdd

	arr.elments = append(arr.elments, values...)

}

func (arr *DynamicArray[T]) RemoveAt(index int) (*T, error) {

	value := &arr.elments[index]

	if index >= arr.len {
		return nil, fmt.Errorf("Index %v out of bounds. len is %v", index, arr.len)
	}
	arr.len -= 1
	arr.elments = append(arr.elments[:index], arr.elments[index+1:]...)

	return value, nil

}
