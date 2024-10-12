package BinaryTree

type BST[T ~int | ~float32] interface {
	isExist(elem T) (error, int) // if exist return index int slice else return error
	Add(elem T)
	Delete(elem T)
}

type Node[T ~int | ~float32] struct {
	val    T
	ind    uint
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

func (n *Node[T]) InitEmptyNode() *Node[T] {
	return &Node[T]{
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (n *Node[T]) InitNodeWithVal(node *Node[T]) *Node[T] {
	return &Node[T]{
		val:    node.val,
		ind:    node.ind,
		left:   node.left,
		right:  node.right,
		parent: node.parent,
	}
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

func (bst *BinarySearchTree[T]) findNode(val T, n *Node[T]) {

}

func (bst *BinarySearchTree[T]) FindMin(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	} else {
		return bst.FindMin(node.left)
	}
}

func (bst *BinarySearchTree[T]) FindMax(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	} else if node.right == nil {
		return node
	} else {
		return bst.FindMax(node.right)
	}
}

func (bst *BinarySearchTree[T]) Add(val T, node *Node[T]) *Node[T] {
	if node != nil {
		if val < node.val {
			node.left = bst.Add(val, node.left)
		} else {
			node.right = bst.Add(val, node.right)
		}
	}
	return node
}

func (bst *BinarySearchTree[T]) Delete(val T, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	} else if val < node.val {
		node.left = bst.Delete(val, node.left)
	} else if val > node.val {
		node.right = bst.Delete(val, node.right)
	} else if node.left != nil && node.right != nil {
		temp := bst.FindMin(node.right)
		node.val = temp.val
		node.right = bst.Delete(node.val, node.right)
	} else {
		temp := node
		if node.left == nil {
			node = node.right
		} else if node.right == nil {

		}
	}
	return node
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
