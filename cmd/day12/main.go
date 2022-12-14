package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

type Point struct {
	x int
	y int
}

type Grid struct {
	grid   [][]int
	width  int
	height int
	start  Point
	end    Point
}

func NewGrid(fileScanner *bufio.Scanner) *Grid {
	grid := Grid{}

	for y := 0; fileScanner.Scan(); y++ {
		row := []int{}
		for x, c := range fileScanner.Text() {
			v := int(c - 0)
			switch v {
			case 83:
				grid.start = Point{x: x, y: y}
				v = 0
			case 69:
				grid.end = Point{x: x, y: y}
				v = 25
			default:
				v -= 97
			}

			row = append(row, v)
		}
		grid.grid = append(grid.grid, row)
	}
	grid.width = len(grid.grid[0])
	grid.height = len(grid.grid)
	return &grid
}

func (grid *Grid) toGraph() *Graph {
	dirs := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	graph := Graph{width: grid.width, height: grid.height}
	graph.coords = make(map[Point]*Node)
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			n := graph.getOrCreateNode(x, y)
			for _, dir := range dirs {
				x2 := x + dir.x
				y2 := y + dir.y
				if x2 >= 0 && y2 >= 0 && x2 < grid.width && y2 < grid.height {
					v := grid.grid[y][x]
					v2 := grid.grid[y2][x2]
					if v2 >= v-1 {
						n2 := graph.getOrCreateNode(x2, y2)
						n.edges = append(n.edges, n2)

					}

				}
			}

		}
	}

	graph.coords[grid.end].cost = 0
	return &graph
}

type Node struct {
	cost    uint
	visited bool
	edges   []*Node
}

func NewNode() *Node {
	node := Node{cost: ^uint(0) - 1}
	node.edges = make([]*Node, 0)
	return &node
}

type Graph struct {
	coords map[Point]*Node
	// Only needed for display
	width  int
	height int
}

func (graph *Graph) getOrCreateNode(x int, y int) *Node {
	p := Point{x: x, y: y}
	n, ok := graph.coords[p]
	if !ok {
		n = NewNode()
		graph.coords[p] = n
	}
	return n
}

func getClosestUnvisited(unvisited []*Node) (*Node, []*Node) {
	cost := ^uint(0)
	index := 0
	for k, v := range unvisited {
		if v.cost < cost {
			cost = v.cost
			index = k
		}
	}
	node := unvisited[index]
	unvisited[index] = unvisited[len(unvisited)-1]
	unvisited[len(unvisited)-1] = nil
	return node, unvisited[:len(unvisited)-1]
}

func (graph *Graph) calcPathCosts() {
	unvisited := []*Node{}
	for _, v := range graph.coords {
		unvisited = append(unvisited, v)
	}
	var node *Node
	for len(unvisited) > 0 {
		node, unvisited = getClosestUnvisited(unvisited)
		for _, n := range node.edges {
			if !n.visited {
				new_cost := node.cost + 1
				if new_cost < n.cost {
					n.cost = new_cost
				}
			}
		}
		node.visited = true
		// fmt.Printf("unvisited: %d\n", len(unvisited))
		// graph.printGraph()
	}
}

func (graph *Graph) printGraph() {
	for y := 0; y < graph.height; y++ {
		for x := 0; x < graph.width; x++ {
			p := Point{x, y}
			if !graph.coords[p].visited {
				fmt.Printf(" ")
			} else if graph.coords[p].cost > 9 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%1d", graph.coords[p].cost)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])
	part1(input)
}

func readInput(filename string) *Grid {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	grid := NewGrid(fileScanner)

	return grid
}

func findMinCostForElevation(graph *Graph, grid *Grid, target_height int) uint {
	cost := ^uint(0)
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if grid.grid[y][x] == target_height {
				p := Point{x, y}
				if graph.coords[p].cost < cost {

					cost = graph.coords[p].cost
				}
			}
		}
	}
	return cost
}

func part1(grid *Grid) {
	graph := grid.toGraph()
	graph.calcPathCosts()
	fmt.Printf("Part 1: %d\n", graph.coords[grid.start].cost)
	fmt.Printf("Part 2: %d\n", findMinCostForElevation(graph, grid, 0))
}
