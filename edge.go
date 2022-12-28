package iGraph

type Edge struct {
	Weight      int
	Origin      CustomNode
	Destination CustomNode
	IsDirected  bool
}
