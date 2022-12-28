package iGraph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	Nodes []CustomNode
	Edges []*Edge
}

func (g *Graph) ReadTxt(path string) {
	// var nodesCount int
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Could not read the file.")
	}

	r := bufio.NewReader(file)
	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			if len(token) == 1 {
				count, err := strconv.Atoi(string(token))
				if err != nil {
					fmt.Println("Unable to convert []byte to int.")
					break
				}
				g.appendNodes(count)

			} else {
				g.insertEdge(string(token))
			}
		}
		if err != nil {
			break
		}
	}

	file.Close()
}

func (g *Graph) insertEdge(line string) {
	a := int(line[0] - '0')
	b := int(line[2] - '0')

	fNode := g.getByValue(a)
	sNode := g.getByValue(b)

	fNode.appendAdj(sNode)
	sNode.appendAdj(fNode)

	newEdge := Edge{
		Origin:      fNode,
		Destination: sNode,
	}
	g.Edges = append(g.Edges, &newEdge)
}

func (g *Graph) appendNodes(count int) {
	for i := 1; i <= count; i++ {
		g.AppendNode(i)
	}
}

func (g *Graph) AppendNode(value int) {
	node := &Node{Value: value}
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) Print() {
	for _, edge := range g.Edges {
		fmt.Print(edge.Origin.GetValue(), "--")
		if edge.IsDirected {
			fmt.Print(">")
		}
		fmt.Print(edge.Destination.GetValue())
	}
	fmt.Println()
}

func (g *Graph) getByValue(value int) *Node {
	for _, node := range g.Nodes {
		if node.GetValue() == value {
			return node.(*Node)
		}
	}
	return nil
}
