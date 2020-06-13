package bintree

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

func traverseNode(node *Node, callback func(int)) {
	if node != nil {
		traverseNode(node.left, callback)
		callback(node.data)
		traverseNode(node.right, callback)
	}
}
