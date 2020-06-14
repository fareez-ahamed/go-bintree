package bintree

import (
	"fmt"
	"strconv"
)

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
func (tree *BinTree) Traverse(callback func(int, int)) {
	if tree == nil || tree.root == nil {
		return
	}
	traverseNode(tree.root, 0, callback)
}

// Print is used for printing the tree
func (tree *BinTree) Print() {
	if tree == nil || tree.root == nil {
		return
	}
	traverseNode(tree.root, 0, func(val int, level int) {
		format := "%" + strconv.Itoa(level*4) + "d\n"
		fmt.Printf(format, val)
	})
}
