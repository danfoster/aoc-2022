package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/danfoster/aoc-2022/internal/common"
	"github.com/danfoster/aoc-2022/internal/graph"
)

type Costs struct {
	Costs  map[string](map[string]*graph.NodeCost)
	Labels []string
	graph  *graph.Graph
}

func (costs *Costs) GetOrCreate(a string, b string) *graph.NodeCost {
	nc, ok := costs.Costs[a][b]
	if !ok {
		nc = graph.NewNodeCost()
		costs.Costs[a][b] = nc
	}
	return nc
}

func (costs *Costs) calcAllPathCosts() {
	for _, v := range costs.Labels {
		costs.calcPathCosts(v)
	}
}

func (costs *Costs) calcPathCosts(starting_label string) {
	starting_nc := costs.GetOrCreate(starting_label, starting_label)
	starting_nc.Cost = 0
	unvisited := *costs.graph.AllNodes()
	var node *graph.Node
	var nc, node_cost *graph.NodeCost
	for len(unvisited) > 0 {
		node, unvisited = costs.getClosestUnvisited(starting_label, unvisited)
		node_cost = costs.Costs[starting_label][node.Label]
		for _, n := range node.Edges {
			nc = costs.GetOrCreate(starting_label, n.Label)
			if !nc.Visited {
				new_cost := node_cost.Cost + 1
				if new_cost < nc.Cost {
					nc.Cost = new_cost
				}
			}
		}
		// fmt.Println(unvisited)
		node_cost.Visited = true
	}
}

func (costs *Costs) getClosestUnvisited(starting_label string, unvisited []*graph.Node) (*graph.Node, []*graph.Node) {
	cost := ^uint(0) - 1
	index := 0
	for k, v := range unvisited {
		n := costs.GetOrCreate(starting_label, v.Label)
		if n.Cost < cost {
			cost = n.Cost
			index = k
		}
	}
	node := unvisited[index]
	unvisited[index] = unvisited[len(unvisited)-1]
	unvisited[len(unvisited)-1] = nil
	return node, unvisited[:len(unvisited)-1]
}

func (costs *Costs) findBestRoute(current_label string, remaining_time int, remaining_nodes []string) int {
	current_pressure := 0
	for _, node := range remaining_nodes {
		nodes := common.RemoveFromStringSlice(remaining_nodes, node)
		time := remaining_time - int(costs.Costs[current_label][node].Cost)
		v := costs.findBestRouteStep(node, nodes, 0, time)
		if v > current_pressure {
			current_pressure = v
		}
	}

	return current_pressure
}

func (costs *Costs) findBestRouteStep(current_label string, remaining_nodes []string, current_pressure int, remaining_time int) int {
	if remaining_time <= 0 {
		return current_pressure
	}

	if costs.graph.Nodes[current_label].Weight != 0 {
		remaining_time--
		pressure := remaining_time * costs.graph.Nodes[current_label].Weight
		current_pressure += pressure

	}
	new_max := current_pressure
	for _, node := range remaining_nodes {
		nodes := common.RemoveFromStringSlice(remaining_nodes, node)
		time := remaining_time - int(costs.Costs[current_label][node].Cost)

		v := costs.findBestRouteStep(node, nodes, current_pressure, time)
		if v > new_max {
			new_max = v
		}
	}

	return new_max
}

func NewCosts(g *graph.Graph, labels []string) *Costs {
	costs := Costs{graph: g}
	size := len(labels)
	costs.Labels = labels
	costs.Costs = make(map[string](map[string]*graph.NodeCost), size)
	for i := 0; i < size; i++ {
		costs.Costs[labels[i]] = make(map[string]*graph.NodeCost, size)
	}
	return &costs
}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	day16(os.Args[1])
}

func day16(filename string) {

	graph := readInput(filename)
	nodes := graph.NodesWithCost()
	*nodes = append(*nodes, graph.Get("AA"))

	node_labels := make([]string, len(*nodes))
	for k, n := range *nodes {
		node_labels[k] = n.Label
	}

	costs := NewCosts(graph, node_labels)
	costs.calcAllPathCosts()

	part1(costs)
	part2(costs, &node_labels)

}

func part1(costs *Costs) {
	result := costs.findBestRoute("AA", 30, common.RemoveFromStringSlice(costs.Labels, "AA"))
	fmt.Println(result)
}

func part2(costs *Costs, node_labels *[]string) {
	pairs := sliceToPairCombos(common.RemoveFromStringSlice(*node_labels, "AA"))
	max := 0
	for _, p := range pairs {
		r1 := costs.findBestRoute("AA", 26, p[0])
		r2 := costs.findBestRoute("AA", 26, p[1])
		sum := r1 + r2
		// fmt.Printf("%v: %d\n", p, sum)
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func sliceToPairCombos(input []string) [][][]string {
	results := [][][]string{}

	size := len(input)

	for i := 0; i < (1 << (size - 1)); i++ {
		if countBits(i) != size/2 {
			continue
		}
		l := []string{}
		r := []string{}
		n := i
		for j := 0; j < size; j++ {
			if n&1 == 0 {
				l = append(l, input[j])
			} else {
				r = append(r, input[j])
			}
			n >>= 1
		}
		combo := [][]string{l, r}
		results = append(results, combo)
	}

	return results
}

func readInput(filename string) *(graph.Graph) {
	readFile, err := os.Open(filename)
	common.Check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	graph := graph.NewGraph()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		re := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? (.*)")
		matches := re.FindStringSubmatch(line)
		label := matches[1]
		weight, err := strconv.Atoi(matches[2])
		common.Check(err)
		edges := strings.Split(matches[3], ", ")
		graph.AddNode(label, weight, edges)

	}

	return graph
}
