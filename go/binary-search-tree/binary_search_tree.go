package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	switch {
	case i > bst.data:
		right := bst.right
		if right == nil {
			bst.right = NewBst(i)
			return
		}
		right.Insert(i)
	default:
		left := bst.left
		if left == nil {
			bst.left = NewBst(i)
			return
		}
		left.Insert(i)
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() (sorted []int) {
	current := bst
	stack := []*BinarySearchTree{}

	for current != nil {
		stack = append(stack, current)
		current = current.left
	}

	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		sorted = append(sorted, node.data)
		sorted = append(sorted, node.right.SortedData()...)
	}
	return sorted
}
