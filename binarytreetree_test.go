package main

import "testing"

type (
	node struct {
		l     *node
		r     *node
		root  *node
		value int
	}
	tree struct {
		name string
		root *node
	}
)

func TestFoundInBinaryTree(t *testing.T) {
	NewTree("tree1", []int{1, 2, 3, 4, 5, 6, 7})
}
