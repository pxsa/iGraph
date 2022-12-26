package iGraph

type Edge struct {
	Weight      int
	Origin      *Node
	Destination *Node
	IsDirected  bool
}
