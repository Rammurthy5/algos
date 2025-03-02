package src

import (
	"fmt"
)

// node represents a single node in the binary tree.
type node struct {
	Value any
	Left  *node
	Right *node
}

// BinaryTree represents the binary tree structure.
type BinaryTree struct {
	Root  *node
	array []*node // Used as a queue for level-order insertion
}

// NewBinaryTree creates an empty binary tree.
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// popleft removes the first element from the array (queue).
func (b *BinaryTree) popleft() {
	if len(b.array) != 0 {
		b.array = b.array[1:] // Shift left (remove first element)
	}
}

// Insert adds a value to the binary tree in level-order (breadth-first).
func (b *BinaryTree) Insert(value any) {
	newNode := &node{Value: value, Left: nil, Right: nil}
	if b.Root == nil {
		b.Root = newNode
		b.array = append(b.array, b.Root)
		return
	}
	// Use array as a queue to insert in level-order
	if b.array[0].Left == nil {
		b.array[0].Left = newNode
		b.array = append(b.array, b.array[0].Left)
		return
	}
	b.array[0].Right = newNode
	b.array = append(b.array, b.array[0].Right)
	b.popleft()
}

// Delete is a placeholder (basic root deletion for now).
func (b *BinaryTree) Delete(value any) {
	if b.Root == nil {
		return
	}
	// Simplest case: Delete root if it matches (for demo purposes)
	if b.Root.Value == value {
		if b.Root.Left == nil {
			b.Root = b.Root.Right
		} else {
			newRoot := b.Root.Left
			if b.Root.Right != nil {
				current := newRoot
				for current.Right != nil {
					current = current.Right
				}
				current.Right = b.Root.Right
			}
			b.Root = newRoot
		}
		// Reset array for consistency
		b.array = []*node{b.Root}
		return
	}
	// TODO: Full deletion logic
	b.popleft()
}

// GetHeight calculates the height of the tree (root is height 0).
func (b *BinaryTree) GetHeight() int {
	return height(b.Root)
}

func height(n *node) int {
	if n == nil {
		return -1
	}
	leftHeight := height(n.Left)
	rightHeight := height(n.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// PreOrderTraversal visits nodes in pre-order (Root, Left, Right).
func (b *BinaryTree) PreOrderTraversal() []any {
	var result []any
	preOrder(b.Root, &result)
	return result
}

func preOrder(n *node, result *[]any) {
	if n != nil {
		*result = append(*result, n.Value)
		preOrder(n.Left, result)
		preOrder(n.Right, result)
	}
}

// InOrderTraversal visits nodes in in-order (Left, Root, Right).
func (b *BinaryTree) InOrderTraversal() []any {
	var result []any
	inOrder(b.Root, &result)
	return result
}

func inOrder(n *node, result *[]any) {
	if n != nil {
		inOrder(n.Left, result)
		*result = append(*result, n.Value)
		inOrder(n.Right, result)
	}
}

// PostOrderTraversal visits nodes in post-order (Left, Right, Root).
func (b *BinaryTree) PostOrderTraversal() []any {
	var result []any
	postOrder(b.Root, &result)
	return result
}

func postOrder(n *node, result *[]any) {
	if n != nil {
		postOrder(n.Left, result)
		postOrder(n.Right, result)
		*result = append(*result, n.Value)
	}
}

// BFSTraversal visits nodes in level-order (breadth-first).
func (b *BinaryTree) BFSTraversal() []any {
	var result []any
	if b.Root == nil {
		return result
	}

	// Use a queue for BFS
	queue := []*node{b.Root}
	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:] // Pop front (shift left)

		// Visit the node
		result = append(result, current.Value)

		// Enqueue children (left first, then right)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return result
}

// Main for testing
func BinaryTreeRun() {
	bt := NewBinaryTree()
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)
	bt.Insert(4)
	bt.Insert(5)
	bt.Insert(6)
	bt.Insert(7)

	fmt.Println("Height:", bt.GetHeight())
	fmt.Println("Pre-Order:", bt.PreOrderTraversal())    // [1 2 4 5 3 6 7]
	fmt.Println("In-Order:", bt.InOrderTraversal())      // [4 2 5 1 6 3 7]
	fmt.Println("Post-Order:", bt.PostOrderTraversal())  // [4 5 2 6 7 3 1]
	fmt.Println("BFS (Level-Order):", bt.BFSTraversal()) // [1 2 3 4 5 6 7]
}
