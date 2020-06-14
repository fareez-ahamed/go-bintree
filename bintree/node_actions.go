package bintree

func addNode(node, newNode *Node) {
	if newNode == nil {
		return
	}
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

func traverseNode(node *Node, level int, callback func(int, int)) {
	if node != nil {
		traverseNode(node.left, level+1, callback)
		callback(node.data, level)
		traverseNode(node.right, level+1, callback)
	}
}
