package iGraph

import (
	"golang.org/x/exp/slices"
)

type Node struct {
	Value int
	Adj   []*Node
}

func (n *Node) appendAdj(node *Node) {
	n.Adj = append(n.Adj, node)
}

func (nn *Node) IsAdjacent(node *Node) bool {
	// check the main node itself
	if nn.Value == node.Value {
		return true
	}
	// check adjacent nodes
	idx := slices.IndexFunc(nn.Adj, func(n *Node) bool {
		return n.Value == node.Value
	})

	// it doesn't find node, so they aren't adjacent
	if idx == -1 {
		return false
	}
	// it's not -1, so it finds the node in its ajd list.
	return true
}

func (n *Node) print() {

}
