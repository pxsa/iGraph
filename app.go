package main

import (
	"pxsa/graph"
)

func main() {
	g := graph.Graph{}
	g.ReadTxt("input.txt")
	g.Print()

	create_independent_graph(&g)
}
