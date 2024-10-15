package BinaryTreeOnGo

import (
	"errors"
	"fmt"
)

type BST[T ~int | ~float32] interface {
	IsExist(elem T) (*Node[T], error)
	Add(elem T) *Node[T]
	Delete(elem T) *Node[T]
}

type Node[T ~int | ~float32] struct {
	val    T
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

func InitEmptyNode[T ~int | ~float32]() *Node[T] {
	return &Node[T]{
		val:    0,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func InitNodeWithVal[T ~int | float32](node *Node[T]) *Node[T] {
	return &Node[T]{
		val:    node.val,
		left:   node.left,
		right:  node.right,
		parent: node.parent,
	}
}

// func (n *Node[T]) delete() {
// 	temp := n.parent
// 	if temp.right.val == n.val {
// 		temp.right = nil
// 	} else {
// 		temp.left = nil
// 	}
// }

// func (n1 *Node[T]) swap(n2 *Node[T]) {
// 	temp := n1.parent
// 	left := n1.left
// 	right := n1.right
// 	val := n1.val
// 	val2 := n2.val
// 	temp2 := n2.parent
// 	n1.left = n2.left
// 	n1.right = n2.right
// 	n1.val = n2.val
// 	if temp.right.val == val {
// 		temp.right = n1
// 	} else {
// 		temp.left = n1
// 	}

// 	n2.left = left
// 	n2.right = right
// 	n2.val = val
// 	if temp2.right.val == val2 {
// 		temp2.right = n2
// 	} else {
// 		temp2.left = n2
// 	}
// }

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

func (bst *BinarySearchTree[T]) Delete(root *Node[T], key T) *Node[T] {
	// nodeToDelete := bst.isExist(val, bst.root)
	// if bst.root == nil || nodeToDelete == nil {
	// 	return nil
	// }
	// if nodeToDelete.right == nil && nodeToDelete.left == nil {
	// 	nodeToDelete.delete()
	// 	return nil
	// }
	// if nodeToDelete.left == nil {
	// 	nodeToDelete.swap(nodeToDelete.right)
	// 	return nodeToDelete
	// }
	// if nodeToDelete.right == nil {
	// 	nodeToDelete.swap(nodeToDelete.left)
	// 	return nodeToDelete
	// }
	// temp := bst.FindMax(nodeToDelete.left)
	// nodeToDelete.val = temp.val
	// temp.delete()
	// return nodeToDelete
	// if bst.root != nil {
	// 	nodeToDelete := bst.isExist(val, bst.root)
	// 	temp := nodeToDelete.parent
	// 	if temp.right.val == nodeToDelete.val {
	// 		temp.right = nil
	// 	} else {
	// 		temp.left = nil
	// 	}
	// }
	// Hz pochemu eto vsya huynya blyat ne rabotaet
	// Rot ebal delete suka
	// }
	// if bst.root != nil {
	// 	nodeToDelete := bst.isExist(val, bst.root)
	// 	temp := nodeToDelete.parent
	// 	if temp.right.val == nodeToDelete.val {
	// 		temp.right = nil
	// 	}
	// 	temp.left = nil
	// }
	if root == nil {
		return nil
	}

	if key < root.val {
		root.left = bst.Delete(root.left, key)
	} else if key > root.val {
		root.right = bst.Delete(root.right, key)
	} else {
		// Node with only one child or no child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		// Node with two children: Get the inorder successor
		minLargerNode := root.right
		for minLargerNode.left != nil {
			minLargerNode = minLargerNode.left
		}

		// Copy the inorder successor's value to this node
		root.val = minLargerNode.val

		// Delete the inorder successor
		root.right = bst.Delete(root.right, minLargerNode.val)
	}

	return root
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
func main() {
	a := InitBST[int]()
	a.Add(1, a.root)
	a.Add(2, a.root)
	a.Add(5, a.root)
	a.Add(4, a.root)
	a.Add(3, a.root)
	_, err := a.IsExist(3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Exist")
	}

	a.Delete(a.root, 3)
	_, err1 := a.IsExist(3)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Exist")
	}

	// Works on bambuk and on usual tree fuuufffff
}
