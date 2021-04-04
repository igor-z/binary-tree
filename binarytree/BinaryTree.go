package binarytree

import "fmt"

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{}
		tree.Root.Value = value

		return
	}

	node := tree.Root
	for {
		if value > node.Value {
			if node.Right == nil {
				node.Right = &Node{
					Value: value,
				}

				break
			} else {
				node = node.Right
			}
		} else if value < node.Value {
			if node.Left == nil {
				node.Left = &Node{
					Value: value,
				}

				break
			} else {
				node = node.Left
			}
		} else {
			break
		}
	}
}

func (tree *BinaryTree) RecursiveInsert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{}
		tree.Root.Value = value
	} else {
		tree.Root.Insert(value)
	}
}

func (tree *BinaryTree) Exists(value int) bool {
	if tree.Root == nil {
		return false
	}

	node := tree.Root
	count := 0
	found := false
	for {
		count++
		if value > node.Value {
			if node.Right == nil {
				break
			} else {
				node = node.Right
			}
		} else if value < node.Value {
			if node.Left == nil {
				break
			} else {
				node = node.Left
			}
		} else {
			found = true
			break
		}
	}

	fmt.Println("it took", count, "iterations")

	return found
}
/*
func (tree *BinaryTree) RecursiveExists(value int) bool {
	if tree.Root == nil {
		return false
	}

	return tree.Root.Exists(value)
}*/