package BinaryTree

import "errors"

type BST[T ~int | ~float32] interface {
	IsExist(elem T) (error, int)
	Add(elem T)
	Delete(elem T)
}

type Node[T ~int | ~float32] struct {
	val   T
	ind   uint
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) InitEmptyNode() *Node[T] {
	return &Node[T]{
		left:  nil,
		right: nil,
	}
}

func (n *Node[T]) InitNodeWithVal(node *Node[T]) *Node[T] {
	return &Node[T]{
		val:   node.val,
		ind:   node.ind,
		left:  node.left,
		right: node.right,
	}
}

func (n *Node[T]) swap(n1 *Node[T], n2 *Node[T]) {
	temp := n1.val
	n1.val = n2.val
	n2.val = temp

	othertemp := n1.ind
	n1.ind = n2.ind
	n2.ind = othertemp
}

type BinarySearchTree[T ~int | ~float32] struct {
	root *Node[T]
}

func (bst *BinarySearchTree[T]) InitBST() *Node[T] {
	return bst.root.InitEmptyNode()

}

func (bst *BinarySearchTree[T]) InitBSTWithVal(rootNode *Node[T]) *Node[T] {
	return bst.root.InitNodeWithVal(rootNode)
}

func (bst *BinarySearchTree[T]) FindMin(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	} else {
		temp := node
		for temp.left != nil {
			temp = temp.left
		}
		return temp
	}
}

func (bst *BinarySearchTree[T]) FindMax(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	} else {
		temp := node
		for temp.right != nil {
			temp = temp.right
		}
		return temp
	}
}

func (bst *BinarySearchTree[T]) Add(val T, node *Node[T]) *Node[T] {
	if node != nil {
		if val < bst.root.val {
			node.left = bst.Add(val, node.left)
		} else {
			node.right = bst.Add(val, node.right)
		}
	}
	return node
}

func (bst *BinarySearchTree[T]) Delete(val T) *Node[T] {
	nodeToDelete := bst.isExist(val, bst.root)
	if bst.root == nil || nodeToDelete == nil {
		return nil
	} else if val < bst.root.val {
		temp := bst.FindMax(nodeToDelete.left)
		bst.root.swap(temp, nodeToDelete)
		nodeToDelete = nil
		return temp
	} else {
		temp := bst.FindMin(nodeToDelete.right)
		bst.root.swap(temp, nodeToDelete)
		nodeToDelete = nil
		return temp
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
	} else {
		return bst.isExist(val, n.left)
	}
}

func (bst *BinarySearchTree[T]) IsExist(val T, n *Node[T]) (*Node[T], error) {
	temp := bst.isExist(val, bst.root)
	if temp != nil {
		return temp, nil
	} else {
		err := errors.New("Node does not exits")
		return nil, err
	}
}
