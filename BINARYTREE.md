A binary tree is a tree data structure in which each node has at most two children,
 referred to as the left child and the right child. it uses Linked List underneath. the difference is LL is
 linear(sequential) and Trees are hierarchical. memory is dynamically allocated unlike array 
 (which is initialized with a fixed size for the continous block of memory for 
 Fast random access O(1) access time using an index).

## types
Full Binary Tree: Every node other than the leaves has two children.
Complete Binary Tree: All levels are fully filled except possibly the last level, which is filled from left to right.
Perfect Binary Tree: All interior nodes have two children and all leaves are at the same level.

## Tree Traversal
Inorder (Left, Root, Right) - useful for sorted BST traversal.
Preorder (Root, Left, Right) - useful for tree reconstruction, copying.
Postorder (Left, Right, Root) - useful for deletion operations.
Level Order - breadth-first traversal
DFS is pre-order, post-order, In-order traversals.

## Usecases
file storage system hierarchies.

## time complexity 
O(n) for all worst case search, insert, delete.