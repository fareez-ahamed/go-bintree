package main

import (
	"fmt"

	"github.com/fareez-ahamed/go-bintree/bintree"
)

func main() {

	var tree bintree.BinTree

	tree.Add(10)
	tree.Add(11)
	tree.Add(9)
	tree.Add(12)

	// fmt.Println(tree)

	tree.Traverse(func(a int) {
		fmt.Println(a)
	})
}
