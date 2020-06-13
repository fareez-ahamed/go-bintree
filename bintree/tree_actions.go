package bintree

// Add is used to add a new value to the Binary tree
func (tree *BinTree) Add(val int) {
	newNode := &Node{data: val}

	if tree.root == nil {
		tree.root = newNode
	} else {
		addNode(tree.root, newNode)
	}

}

// Traverse is used for DFS traverse
func (tree *BinTree) Traverse(callback func(int)) {
	if tree == nil || tree.root == nil {
		return
	}
	traverseNode(tree.root, callback)
}
