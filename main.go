package main

import (
	"fmt"

	"graph/node"
	"graph/ordered"
)

func main() {
	//var root *node.Node
	//root = node.Add(root, 5)
	//root = node.Add(root, 2)
	//root = node.Add(root, 1)
	//root = node.Add(root, 7)
	//root = node.Add(root, 4)
	//root = node.Add(root, 10)
	//
	//root.Print(root)

	tree := node.New(ordered.NewOrderedInt())
	tree.Add(5)
	tree.Add(4)
	tree.Add(6)
	tree.Add(10)
	tree.Add(1)

	tree.Print()

	tree.Delete(5)
	fmt.Println()

	tree.Print()
}
