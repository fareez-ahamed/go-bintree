package main

import (
	"fmt"
	"math/rand"

	"github.com/fareez-ahamed/go-bintree/bintree"
)

func main() {

	var tree bintree.BinTree

	for i := 0; i < 100; i++ {
		tree.Add(rand.Intn(1000))
	}

	tree.Traverse(func(a int) {
		fmt.Println(a)
	})
}
