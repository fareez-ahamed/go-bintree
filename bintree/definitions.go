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
