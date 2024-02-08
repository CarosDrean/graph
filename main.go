package main

import (
	"fmt"

	"graph/node"
	"graph/ordered"
)

func main() {
	tree := node.NewTree(ordered.NewOrderedInt())
	tree.Add(15)
	tree.Add(5)
	tree.Add(4)
	tree.Add(2)
	tree.Add(20)
	tree.Add(11)
	tree.Add(9)
	tree.Add(25)
	tree.Add(35)
	tree.Add(7)
	tree.Add(6)
	tree.Add(3)
	tree.Add(10)
	tree.Add(1)
	tree.Add(16)
	tree.Add(21)
	tree.Add(22)
	tree.Add(23)

	tree.PrintTree()

	fmt.Printf("\n=========================\n")

	tree.Delete(5)
	tree.Delete(1)
	tree.Delete(2)
	tree.Delete(4)
	tree.Delete(21)

	tree.PrintTree()
}
