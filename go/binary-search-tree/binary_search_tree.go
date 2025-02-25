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
		insert(
			bst,
			i,
			func(b *BinarySearchTree) *BinarySearchTree { return b.right },
			func(b *BinarySearchTree, c *BinarySearchTree) { b.right = c },
		)
	default:
		insert(
			bst,
			i,
			func(b *BinarySearchTree) *BinarySearchTree { return b.left },
			func(b *BinarySearchTree, c *BinarySearchTree) { b.left = c },
		)
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

func insert(
	bst *BinarySearchTree,
	i int,
	get func(*BinarySearchTree) *BinarySearchTree,
	set func(*BinarySearchTree, *BinarySearchTree),
) {
	node := get(bst)
	if node == nil {
		set(bst, NewBst(i))
		return
	}
	node.Insert(i)
}
