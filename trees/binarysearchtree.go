package trees

import (
	"cmp"
	"fmt"
	"math"
)

type Node[T cmp.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	root      *Node[T]
	nodeCount int
}

func New[T cmp.Ordered]() *BinarySearchTree[T] {
	t := &BinarySearchTree[T]{}
	t.root = nil
	return t
}

func (t *BinarySearchTree[T]) Empty() bool {
	return t.Size() == 0
}

func (t *BinarySearchTree[T]) Size() int {
	return t.nodeCount
}

func (t *BinarySearchTree[T]) addNode(node *Node[T], v T) *Node[T] {
	if node == nil {
		node = &Node[T]{value: v}
	} else {
		if v < node.value {
			node.left = t.addNode(node.left, v)
		} else {
			node.right = t.addNode(node.right, v)
		}
	}
	t.nodeCount++
	return node
}

func (t *BinarySearchTree[T]) Add(v T) error {
	if t.Contains(v) {
		return fmt.Errorf("Value alreadye exists inside the tree.")
	}
	t.root = t.addNode(t.root, v)
	return nil
}

func (t *BinarySearchTree[T]) Contains(v T) bool {
	return t.contains(t.root, v)
}

func (t *BinarySearchTree[T]) contains(node *Node[T], v T) bool {

	if node == nil {
		return false
	}

	if node.value > v {
		return t.contains(node.left, v)
	} else if node.value < v {
		return t.contains(node.right, v)
	}
	return true
}

func (t *BinarySearchTree[T]) Remove(v T) error {
	if !t.Contains(v) {
		return fmt.Errorf("Value doesn't exist inside the tree. ")
	}

	t.root = t.removeNode(t.root, v)
	return nil
}

func (t *BinarySearchTree[T]) removeNode(node *Node[T], v T) *Node[T] {
	if node == nil {
		return nil
	}

	if node.value > v {
		node.left = t.removeNode(node.left, v)
	} else if node.value < v {
		node.right = t.removeNode(node.right, v)
	} else {
		if node.left == nil {
			rightChild := node.right
			node = nil
			t.nodeCount--
			return rightChild
		} else if node.right == nil {
			leftChild := node.left
			node = nil
			t.nodeCount--
			return leftChild
		} else {
			// Find the smallest element in the right sub tree and replace the node with that
			tmp := t.digLeft(node.right)
			node.value = tmp
			node.right = t.removeNode(node.right, tmp)
		}

	}
	return node
}

func (t *BinarySearchTree[T]) digLeft(node *Node[T]) T {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr.value
}

func (t *BinarySearchTree[T]) Height() int {
	return t.computeHeight(t.root)
}

func (t *BinarySearchTree[T]) computeHeight(node *Node[T]) int {
	if node == nil {
		return 0
	}

	leftSubTreeHeight := t.computeHeight(node.left)
	rightSubTreeHeight := t.computeHeight(node.right)

	max := math.Max(float64(leftSubTreeHeight), float64(rightSubTreeHeight))
	return int(max) + 1
}
