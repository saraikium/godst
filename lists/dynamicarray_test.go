package dynamicarray

import (
	"testing"
)

func TestDynamicArrayNew(t *testing.T) {
	list := New[int](10)

	if is_empty := list.Empty(); !is_empty {
		t.Errorf("Expected list to be empty at creation. Expected list.Empty() = true, got %v", is_empty)
	}

	if size := list.Size(); size != 0 {
		t.Errorf("Expected list to be empty at creation. expected list.Size() = 0, got %v", size)
	}

}
