package binarytree

type Node struct {
	Left *Node
	Right *Node
	Value int
}

func (node *Node) Insert(value int) {
	if value > node.Value {
		if node.Right == nil {
			node.Right = &Node{
				Value: value,
			}
		} else {
			node.Right.Insert(value)
		}
	} else if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{
				Value: value,
			}
		} else {
			node.Left.Insert(value)
		}
	}
}

/*func (node *Node) Exists(value int) bool {

}*/