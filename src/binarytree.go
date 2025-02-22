package src

// has a root node and then each node has a left and right child. last node should point to nil.
// get height
// do 4 different traversals
type node struct {
	Value any
	Left  *node
	Right *node
}

type BinaryTree struct {
	Root  *node
	array []*node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) pop() {
	if len(b.array) != 0 {
		b.array = b.array[:len(b.array)-1]
	}
}

func (b *BinaryTree) Insert(value any) {
	if b.Root == nil {
		b.Root = &node{Value: value, Left: nil, Right: nil}
		b.array = append(b.array, b.Root)
		return
	}
	if b.array[0].Left == nil {
		b.array[0].Left = &node{Value: value, Left: nil, Right: nil}
		b.array = append(b.array, b.array[0].Left)
		return
	}
	b.array[0].Right = &node{Value: value, Left: nil, Right: nil}
	b.array = append(b.array, b.array[0].Right)
	b.pop()
}

func (b *BinaryTree) Delete(value any) {
	if b.Root != nil {
		// ToDo

	}
	b.pop()
}

func (b *BinaryTree) InOrderTraversal() {
	if b.Root != nil {
		// ToDo
	}
}
