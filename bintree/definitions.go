package bintree

// Node is the basic building block of the
// Binary tree.
type Node struct {
	data  int
	left  *Node
	right *Node
}

// BinTree is the Binary Tree datastructure
type BinTree struct {
	root *Node
}

// Add is used to add a new value to the Binary tree
func (tree *BinTree) Add(val int) {
	newNode := &Node{data: val}

	if tree.root == nil {
		tree.root = newNode
	} else {
		addNode(tree.root, newNode)
	}

}

func addNode(node *Node, newNode *Node) {
	if node.data < newNode.data {
		if node.right == nil {
			node.right = newNode
			return
		}
		addNode(node.right, newNode)
	} else {
		if node.left == nil {
			node.left = newNode
			return
		}
		addNode(node.left, newNode)
	}
}

// Traverse is used for DFS traverse
func (tree *BinTree) Traverse(callback func(int)) {
	if tree == nil || tree.root == nil {
		return
	}
	traverseNode(tree.root, callback)
}

func traverseNode(node *Node, callback func(int)) {
	if node != nil {
		traverseNode(node.left, callback)
		callback(node.data)
		traverseNode(node.right, callback)
	}
}
