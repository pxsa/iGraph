package graph

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Node struct {
	Value    int
	adj      []*Node
	Parent   *Node
	children []*Node
}

func (n *Node) appendAdj(node *Node) {
	n.adj = append(n.adj, node)
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) InsertChild(child *Node) {
	//BUG
	n.children = append(n.children, child)
	// if 
	if !n.isChildOf(child) {
		if child.Parent == nil {

			child.Parent = n
		}
	}
}

func (nn *Node) IsAdjacent(node *Node) bool {
	// check the main node itself
	if nn.Value == node.Value {
		return true
	}
	// check adjacent nodes
	idx := slices.IndexFunc(nn.adj, func(n *Node) bool {
		return n.Value == node.Value
	})

	// it doesn't find node, so they aren't adjacent
	if idx == -1 {
		return false
	}
	// it's not -1, so it finds the node in its ajd list.
	return true
}


func (n *Node) isChildOf(node *Node) bool{
	// check children nodes
	idx := slices.IndexFunc(node.children, func(child *Node) bool {
		return child == n
	})

	if idx == -1 {
		return false
	}

	return true
}


func (n *Node) print() {
	for _, child := range n.children {
		fmt.Printf("%d ",child.Value)
		child.print()
	}
	fmt.Print(", ")
	
}