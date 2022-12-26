# iGraph
iGraph is a package written in go which you can easily create different types of graph such as tree data structure along side with some more facilities such as inserting vertices and edges, deleteing them and much more.

## Content
- [Installation](#Installation)
- [Node](#Node)
  - [Simple usage](#Simple-usage)
- [Graph](#Graph)

## Installation
To install `iGraph` package, you need to install Go and set your Go wrokspace first.
1. You first need [Go](https://go.dev) installed, then you can use the below Go command to install `iGraph`.
```
go get -u github.com/pxsa/iGraph
```
2. Then you can simpily import `iGraph` package and have fun;)

## Node
`Node` is nothing more than a type that you have to create and insert to your `Graph` object. It has some useful attributes and methods which are listed below.

### Simple usage
```
myNode := iGraph.Node {
  Value: 1,
  Parent: nil,
}
```

## Graph
graph is a data structure that has two important attributes `Nodes` and `Edges`. so here comes the question that how we can initialize Graph. for this trouble we can use `Read_txt(path string)` func, which only takes one parameter as the location of the file.

> please pay attention that the first line of your file should contains a number to indicate the count's of nodes in your graph.

see [input](https://github.com/pxsa/iGraph/blob/master/input.txt) for more informaition.
### Simple usage
```
myGraph := graph.Graph{}
g.ReadTxt("input.txt")
```
