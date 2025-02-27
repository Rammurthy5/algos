package src

import "errors"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type BinarySearchTree struct {
	root *bstnode
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(item int) {
	newNode := &bstnode{value: item, left: nil, right: nil}
	if bst.root == nil {
		bst.root = newNode
		return
	}
	current := bst.root
	for current != nil {
		if item < current.value {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		}
	}
}

func (bst *BinarySearchTree) Search(item int) *bstnode {
	if bst.root == nil {
		return nil
	}
	current := bst.root
	for current != nil {
		if item == current.value {
			return current
		}
		if item < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func (bst *BinarySearchTree) Delete(item int) (bool, error) {
	var parent *bstnode
	current := bst.root
	for current != nil && current.value != item {
		parent = current
		if item < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	if current == nil {
		return false, errors.New("Node not found")
	}

	// Case 1: No children
	if current.left == nil && current.right == nil {
		if current == bst.root {
			bst.root = nil
		} else if parent.left == current {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return true, nil
	}

	// Case 2: One child
	if current.left == nil {
		if current == bst.root {
			bst.root = current.right
		} else if parent.left == current {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
		return true, nil
	}
	if current.right == nil {
		if current == bst.root {
			bst.root = current.left
		} else if parent.left == current {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
		return true, nil
	}

	// Case 3: Two children
	successorParent := current
	successor := current.right
	for successor.left != nil {
		successorParent = successor
		successor = successor.left
	}
	current.value = successor.value
	if successorParent.left == successor {
		successorParent.left = successor.right
	} else {
		successorParent.right = successor.right
	}
	return true, nil
}
