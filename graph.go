package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	Nodes []*Node
	Edges [][2]int
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

	slice := [2]int{a, b}
	g.Edges = append(g.Edges, slice)
}

func (g *Graph) appendNodes(count int) {
	for i:=1; i <= count; i++{
		g.AppendNode(i)
	}
}

func (g *Graph) AppendNode(value int) {
	node := &Node{Value: value, Parent: nil}
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) Print() {
	for _, edge := range g.Edges {
		fmt.Println(edge)
	}
}

func (g *Graph) getByValue(value int) *Node{
	for _, node := range(g.Nodes) {
		if node.Value == value {
			return node
		}
	}
	return nil
} 

