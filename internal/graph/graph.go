package graph

import "fmt"

type Node struct {
	Label  string
	Weight int
	Edges  []*Node
}

type Graph struct {
	Nodes map[string]*Node
}

func (graph *Graph) AddNode(label string, weight int, links []string) *Node {
	node := graph.GetOrCreate(label)
	node.Weight = weight
	for _, link := range links {
		n := graph.GetOrCreate(link)
		node.Edges = append(node.Edges, n)
	}
	return node
}

func (graph *Graph) GetOrCreate(label string) *Node {
	node, ok := graph.Nodes[label]
	if !ok {
		node = &Node{Label: label, Weight: 99}
		graph.Nodes[label] = node
	}
	return node
}

func (graph *Graph) Get(label string) *Node {
	node, ok := graph.Nodes[label]
	if !ok {
		panic("Node doesn't exist")
	}
	return node
}

func (graph *Graph) ToDot() {
	fmt.Printf("graph G {\n")
	for _, n := range graph.Nodes {
		fmt.Printf("  %s [label=\"%s [%d]\"]\n", n.Label, n.Label, n.Weight)
		for _, e := range n.Edges {
			fmt.Printf("  %s -- %s\n", n.Label, e.Label)
		}
	}
	fmt.Printf("}\n")
}

func (graph *Graph) NodesWithCost() *[](*Node) {
	nodes := []*Node{}
	for _, n := range graph.Nodes {
		if n.Weight > 0 {
			nodes = append(nodes, n)
		}
	}
	return &nodes
}

func (graph *Graph) AllNodes() *[](*Node) {
	nodes := []*Node{}
	for _, n := range graph.Nodes {
		nodes = append(nodes, n)
	}
	return &nodes
}

func NewGraph() *Graph {
	graph := Graph{}
	graph.Nodes = make(map[string]*Node)
	return &graph
}

type NodeCost struct {
	Visited bool
	Cost    uint
}

func NewNodeCost() *NodeCost {
	nc := NodeCost{Cost: ^uint(0) - 1}
	return &nc
}
