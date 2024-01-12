package lists

import (
	"github.com/saraikium/godst/containers"
	"github.com/saraikium/godst/utils"
)

type List[T any] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(value T) bool
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	containers.Container[T]
}
