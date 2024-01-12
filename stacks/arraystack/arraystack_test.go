package arraystack

import "testing"

func TestStackPush(t *testing.T) {
	stack := New[int](10)
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
}
