package lists

import (
	"github.com/saraikium/utils"
)

type List[T any] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
