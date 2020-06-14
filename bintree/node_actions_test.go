package bintree

import "testing"

func TestAddNode(t *testing.T) {
	node := &Node{data: 1}
	newNode := &Node{data: 2}
	addNode(node, newNode)
	if node.right != newNode {
		t.Error("Adding node greater than current node is not working")
	}

	newNode = &Node{data: 0}
	addNode(node, newNode)
	if node.left != newNode {
		t.Error("Adding node lesser than current node is not working")
	}

	node = &Node{data: 1}
	newNode = nil
	addNode(node, newNode)
	if node.left != nil || node.right != nil {
		t.Error("Adding nil node is not working propertly")
	}

	nodeRight := &Node{data: 20}
	nodeLeft := &Node{data: 10}
	node = &Node{data: 15, left: nodeLeft, right: nodeRight}
	newNode = &Node{data: 17}
	addNode(node, newNode)
	if nodeRight.left != newNode {
		t.Error("Adding second level on the right side")
	}

	newNode = &Node{data: 13}
	addNode(node, newNode)
	if nodeLeft.right != newNode {
		t.Error("Adding second level on the left side")
	}
}
