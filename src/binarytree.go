package src
import "slices"

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

func (b *BinaryTree) popleft() {
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
	b.popleft()
}

func pop(temp []*node) []*node {
	if len(temp) == 0 {
		return temp
	}
	return temp[:len(temp)-1]
}

func inorder_helper(temp []*node, visited []int) {
	// change the line 65 and other conditions only to check visited rather .left or .right nil
	if len(temp) == 0 {
		return
	}
	last_ind := len(temp) - 1
	if temp[last_ind].Left != nil {
		if !slices.Contains(visited, temp[last_ind].Left.Value) {	
			temp = append(temp, temp[last_ind].Left)
			inorder_helper(temp)
		}
		inorder_helper(temp)
	} else {
		if !slices.Contains(visited, temp[last_ind].Value) {
			visited = append(visited, temp[last_ind].Value)
			temp[last_ind].Left = nil
			
			if temp[last_ind].Right != nil {
				newLast := temp[last_ind].Right
				temp[last_ind-1].Left = nil
				pop(temp)
				temp = append(temp, newLast)
				inorder_helper(temp)
			} else {
				temp[last_ind].Right = nil
			}
		}
		pop(temp)
		inorder_helper(temp)
	}
}
1,2,3,4,5,6,7,8
3,
4,8,2,1,6,

    1
 2	 3
4 5 6 7
 8	

func (b *BinaryTree) InOrderTraversal() {
	if b.Root != nil {
		inorder_helper([]*node{b.Root})
	}
}
