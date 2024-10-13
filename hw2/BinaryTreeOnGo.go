package BinaryTreeOnGo

import (
	"errors"
)

type BST[T ~int | ~float32] interface {
	IsExist(elem T) (*Node[T], error)
	Add(elem T) *Node[T]
	Delete(elem T) *Node[T]
}

type Node[T ~int | ~float32] struct {
	val    T
	ind    uint
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

func InitEmptyNode[T ~int | ~float32]() *Node[T] {
	return &Node[T]{
		val:    0,
		ind:    0,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func InitNodeWithVal[T ~int | float32](node *Node[T]) *Node[T] {
	return &Node[T]{
		val:    node.val,
		ind:    node.ind,
		left:   node.left,
		right:  node.right,
		parent: node.parent,
	}
}

// func (n *Node[T]) swap(n1 *Node[T], n2 *Node[T]) {
// 	temp := n1.val
// 	n1.val = n2.val
// 	n2.val = temp

// 	othertemp := n1.ind
// 	n1.ind = n2.ind
// 	n2.ind = othertemp
// }

type BinarySearchTree[T ~int | ~float32] struct {
	root *Node[T]
}

func InitBST[T ~int | ~float32]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root: InitEmptyNode[T](),
	}
}

func InitBSTWithVal[T ~int | ~float32](rootNode *Node[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root: rootNode,
	}
}

func (bst *BinarySearchTree[T]) FindMin(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	temp := node
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func (bst *BinarySearchTree[T]) FindMax(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	temp := node
	for temp.right != nil {
		temp = temp.right
	}
	return temp
}

func (bst *BinarySearchTree[T]) Add(val T, node *Node[T]) *Node[T] {
	if node != nil {
		if val < node.val {
			node.left = bst.Add(val, node.left)
			node.left.parent = node
		} else {
			node.right = bst.Add(val, node.right)
			node.right.parent = node
		}
	} else {
		node = &Node[T]{
			val:    val,
			parent: nil,
		}
	}
	return node
}

func (bst *BinarySearchTree[T]) Delete(val T) {
	if bst.root != nil {
		nodeToDelete := bst.isExist(val, bst.root)
		temp := nodeToDelete.parent
		if temp.right.val == nodeToDelete.val {
			temp.right = nil
		}
		temp.left = nil
	}
}

func (bst *BinarySearchTree[T]) isExist(val T, n *Node[T]) *Node[T] {
	if n == nil {
		return nil
	}
	if n.val == val {
		return n
	}

	if n.val < val {
		return bst.isExist(val, n.right)
	}
	return bst.isExist(val, n.left)
}

func (bst *BinarySearchTree[T]) IsExist(val T) (*Node[T], error) {
	temp := bst.isExist(val, bst.root)
	if temp != nil {
		return temp, nil
	}
	err := errors.New("Node does not exits")
	return nil, err
}
