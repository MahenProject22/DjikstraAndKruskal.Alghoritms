package main

import (
	"fmt"
	"math"
)

// GraphEdge represents a single undirected graph edge.
type GraphEdge struct {
	From   int
	To     int
	Weight int
}

// Graph stores the adjacency list of the graph.
type Graph struct {
	adj map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int]map[int]int)}
}

// AddEdge adds an undirected edge to the graph.
func (g *Graph) AddEdge(u, v, weight int) {
	if g.adj[u] == nil {
		g.adj[u] = make(map[int]int)
	}
	if g.adj[v] == nil {
		g.adj[v] = make(map[int]int)
	}
	g.adj[u][v] = weight
	g.adj[v][u] = weight
}

// Nodes returns every node of the graph.
func (g *Graph) Nodes() []int {
	nodes := make([]int, 0, len(g.adj))
	for n := range g.adj {
		nodes = append(nodes, n)
	}
	return nodes
}

// Dijkstra returns the shortest distance to every node and the previous
// node on each shortest path. It works only on graphs with positive weights.
func Dijkstra(g *Graph, start int) (map[int]int, map[int]int) {
	dist := make(map[int]int)
	prev := make(map[int]int)
	visited := make(map[int]bool)

	for _, n := range g.Nodes() {
		dist[n] = math.MaxInt
	}
	dist[start] = 0

	for {
		// Pick the unvisited node with the smallest distance.
		current := -1
		for _, n := range g.Nodes() {
			if visited[n] {
				continue
			}
			if current == -1 || dist[n] < dist[current] {
				current = n
			}
		}

		// Stop when no reachable node is left.
		if current == -1 || dist[current] == math.MaxInt {
			break
		}
		visited[current] = true

		// Relax every edge leaving the current node.
		for neighbor, weight := range g.adj[current] {
			candidate := dist[current] + weight
			if candidate < dist[neighbor] {
				dist[neighbor] = candidate
				prev[neighbor] = current
			}
		}
	}

	return dist, prev
}

// Path rebuilds the shortest path from start to target using prev.
func Path(prev map[int]int, start, target int) []int {
	if _, ok := prev[target]; !ok {
		if target == start {
			return []int{start}
		}
		return nil
	}

	path := []int{target}
	for path[len(path)-1] != start {
		path = append(path, prev[path[len(path)-1]])
	}

	// Reverse the path so it starts from the source node.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 4)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 7)
	g.AddEdge(4, 5, 3)

	start := 1

	dist, prev := Dijkstra(g, start)

	fmt.Printf("Shortest paths from node %d:\n", start)
	for _, node := range g.Nodes() {
		if dist[node] == math.MaxInt {
			fmt.Printf("  %d -> %d: unreachable\n", start, node)
			continue
		}
		fmt.Printf("  %d -> %d: cost %d  path %v\n", start, node, dist[node], Path(prev, start, node))
	}
}
