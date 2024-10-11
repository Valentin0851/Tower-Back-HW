package BinaryTree

type BST[T any] interface {
	isExist(elem T) (error, int) // if exist return index int slice else return error
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
	return node
}
