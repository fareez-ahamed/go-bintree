package main

import (
	"math/rand"

	"github.com/fareez-ahamed/go-bintree/bintree"
)

func main() {

	var tree bintree.BinTree

	for i := 0; i < 20; i++ {
		tree.Add(rand.Intn(1000))
	}

	tree.Print()

	// tree.Traverse(func(a int, level int) {
	// 	fmt.Println(a)
	// })
}
