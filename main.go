package main

import (
	"fake.com/binarytree"
	"fmt"
	"strconv"
)

func main() {
	tree := &binarytree.BinaryTree{}
	values := [...]int{
		100,
		202,
		500,
		87,
		99,
		1265,
		87,
	}

	for _, v := range values {
		tree.Insert(v)
	}

	printTree(tree)

	fmt.Println("10 must be false", tree.Exists(10))
	fmt.Println("87 must be true", tree.Exists(87))
	fmt.Println("1265 must be true", tree.Exists(1265))
}

func printTree(tree *binarytree.BinaryTree) {
	var nodes []*binarytree.Node

	nodes = append(nodes, tree.Root)

	for len(nodes) > 0 {
		nextLevelNodes := make([]*binarytree.Node, 0, len(nodes) * 2)

		outputStr := ""
		notNilNodesCount := 0
		for i, node := range nodes {
			if i > 0 && i % (len(nodes) / 2) == 0 {
				outputStr += "| "
			}

			if node == nil {
				outputStr += "- "

				nextLevelNodes = append(nextLevelNodes, nil)
				nextLevelNodes = append(nextLevelNodes, nil)
			} else {
				notNilNodesCount++
				outputStr += strconv.FormatInt(int64(node.Value), 10) + " "

				nextLevelNodes = append(nextLevelNodes, node.Left)
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		if notNilNodesCount == 0 {
			break
		}

		fmt.Println(outputStr)

		nodes = nextLevelNodes
	}
}