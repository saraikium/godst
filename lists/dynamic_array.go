package dynamicarray

import (
	"golang.org/x/exp/constraints"
)

type DynamicArray[T any] struct {
	elments  []T
	len      uint64
	capacity uint64
}



