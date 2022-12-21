package graph

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
	n.children = append(n.children, child)
}