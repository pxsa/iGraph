package main

import "pxsa/graph"

func create_independent_graph(iGraph *graph.Graph) {
	
	iTree := graph.Tree {
		Root: &graph.Node{
			Value: 0,
			Parent: nil,
		},
	}

	recursive(iTree.Root, iGraph)
	
}

func recursive(root *graph.Node, iGraph *graph.Graph) {
	node := iGraph.Nodes[0]
	
	root.InsertChild(node)

	recursive(root, iGraph)
}