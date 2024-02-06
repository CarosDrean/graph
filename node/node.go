package node

import (
	"fmt"

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

func New(ordered ordered.Ordered) Tree {
	return Tree{ordered: ordered}
}

func (t *Tree) Add(value any) {
	t.root = t.add(t.root, value)
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

	//balance := calculateBalance(current)

	return current
}

func (t *Tree) Delete(data any) {
	t.root = t.delete(t.root, data)
}

func (t *Tree) delete(n *Node, data any) *Node {
	if n == nil {
		return n
	}

	if t.ordered.IsEqual(n.data, data) {
		// ifs replace actual node whit opposite node, opposite node can be nil
		if n.left == nil {
			return n.right
		}

		if n.right == nil {
			return n.left
		}

		// here use right for no break the tree
		minRight := t.findMin(n.right)
		n.data = minRight.data

		n.right = t.delete(n.right, minRight.data)

		return n
	}

	if t.ordered.IsLeft(n.data, data) {
		return t.delete(n.left, data)
	}

	return t.delete(n.right, data)
}

func (t *Tree) findMin(n *Node) *Node {
	if n.left == nil {
		return n
	}

	return t.findMin(n.left)
}

func (t *Tree) Print() {
	t.print(t.root)
}

func (t *Tree) print(n *Node) {
	if n != nil {
		t.print(n.left)
		fmt.Printf("%v \n", n.data)
		t.print(n.right)
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
