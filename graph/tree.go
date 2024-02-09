package graph

import (
	"fmt"
	"log"

	"graph/ordered"
)

type node struct {
	data   any
	height int
	left   *node
	right  *node
}

type Tree struct {
	root *node

	ordered ordered.Ordered
}

func NewTree(ordered ordered.Ordered) Tree {
	return Tree{ordered: ordered}
}

func (t *Tree) Add(value any) {
	if !t.ordered.IsValidType(value) {
		log.Printf("the type of %v is invalid\n", value)
		return
	}

	t.root = t.add(t.root, value)
}

func (t *Tree) Delete(value any) {
	if !t.ordered.IsValidType(value) {
		log.Printf("the type of %v is invalid\n", value)
		return
	}

	t.root = t.delete(t.root, value)
}

func (t *Tree) PrintTree() {
	printTree(t.root, 0)
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
}

func (t *Tree) add(current *node, data any) *node {
	if current == nil {
		return &node{data: data, height: 1}
	}

	if t.ordered.IsEqual(current.data, data) {
		// no add duplicated data
		return current
	}

	if t.ordered.IsLeft(current.data, data) {
		current.left = t.add(current.left, data)
	} else {
		current.right = t.add(current.right, data)
	}

	// one is for the current graph
	current.height = 1 + max(height(current.left), height(current.right))

	current = balancer(current)

	return current
}

func (t *Tree) delete(current *node, data any) *node {
	if current == nil {
		return nil
	}

	if t.ordered.IsEqual(current.data, data) {
		// this ifs replace actual graph whit opposite graph, opposite graph can be nil
		if current.left == nil {
			return current.right
		}

		if current.right == nil {
			return current.left
		}

		// find the node with the lowest value of the largest subtree so that it remains as a replacement
		// for the node that is being eliminated, then, since it is a replacement for the node that is being eliminated,
		// eliminate the node from the right subtree
		minRight := findMin(current.right)
		current.data = minRight.data

		current.right = t.delete(current.right, minRight.data)
	}

	if t.ordered.IsLeft(current.data, data) {
		current.left = t.delete(current.left, data)
	} else {
		current.right = t.delete(current.right, data)
	}

	current.height = 1 + max(height(current.left), height(current.right))

	current = balancer(current)

	return current
}

func balancer(node *node) *node {
	balance := calculateBalance(node)

	// left - left
	if balance > 1 && calculateBalance(node.left) >= 0 {
		return rotateRight(node)
	}

	// left - right
	if balance > 1 && calculateBalance(node.left) < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	// right - right
	if balance < -1 && calculateBalance(node.right) <= 0 {
		return rotateLeft(node)
	}

	// right - left
	if balance < -1 && calculateBalance(node.right) > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	return node
}

func rotateRight(node *node) *node {
	childLeft := node.left
	childRightOfLeft := childLeft.right

	childLeft.right = node
	node.left = childRightOfLeft

	node.height = max(height(node.left), height(node.right)) + 1
	childLeft.height = max(height(childLeft.left), height(childLeft.right)) + 1

	return childLeft
}

func rotateLeft(node *node) *node {
	childRight := node.right
	childLeftOfRight := childRight.left

	childRight.left = node
	node.right = childLeftOfRight

	node.height = max(height(node.left), height(node.right)) + 1
	childRight.height = max(height(childRight.left), height(childRight.right)) + 1

	return childRight
}

func calculateBalance(n *node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func height(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func findMin(n *node) *node {
	if n.left == nil {
		return n
	}

	return findMin(n.left)
}

func printTree(root *node, spacing int) {
	if root == nil {
		return
	}

	spacing += 5

	printTree(root.right, spacing)

	fmt.Println()
	for i := 5; i < spacing; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d", root.data)

	printTree(root.left, spacing)
}

func printInOrder(n *node) {
	if n != nil {
		printInOrder(n.left)
		fmt.Printf("%v \n", n.data)
		printInOrder(n.right)
	}
}
