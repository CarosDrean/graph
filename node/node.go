package node

import (
	"fmt"
	"log"

	"graph/ordered"
)

type Node struct {
	data   any
	height int
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node

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
	t.printTree(t.root, 0)
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
}

func (t *Tree) printTree(root *Node, spacing int) {
	if root == nil {
		return
	}

	spacing += 5

	t.printTree(root.right, spacing)

	fmt.Println()
	for i := 5; i < spacing; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d", root.data)

	t.printTree(root.left, spacing)
}

func (t *Tree) add(current *Node, data any) *Node {
	if current == nil {
		return &Node{data: data, height: 1}
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

	// one is for the current node
	current.height = 1 + max(height(current.left), height(current.right))

	current = balancer(current)

	return current
}

func (t *Tree) delete(current *Node, data any) *Node {
	if current == nil {
		return nil
	}

	if t.ordered.IsEqual(current.data, data) {
		// ifs replace actual node whit opposite node, opposite node can be nil
		if current.left == nil {
			return current.right
		}

		if current.right == nil {
			return current.left
		}

		// here use right for no break the tree
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

func balancer(node *Node) *Node {
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

func rotateRight(node *Node) *Node {
	childLeft := node.left
	childRightOfLeft := childLeft.right

	childLeft.right = node
	node.left = childRightOfLeft

	node.height = max(height(node.left), height(node.right)) + 1
	childLeft.height = max(height(childLeft.left), height(childLeft.right)) + 1

	return childLeft
}

func rotateLeft(node *Node) *Node {
	childRight := node.right
	childLeftOfRight := childRight.left

	childRight.left = node
	node.right = childLeftOfRight

	node.height = max(height(node.left), height(node.right)) + 1
	childRight.height = max(height(childRight.left), height(childRight.right)) + 1

	return childRight
}

func findMin(n *Node) *Node {
	if n.left == nil {
		return n
	}

	return findMin(n.left)
}

func printInOrder(n *Node) {
	if n != nil {
		printInOrder(n.left)
		fmt.Printf("%v \n", n.data)
		printInOrder(n.right)
	}
}

func calculateBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}
