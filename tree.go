package iGraph

type Tree struct {
	Root *Node
}

func (t *Tree) Print() {
	t.Root.print()
}
