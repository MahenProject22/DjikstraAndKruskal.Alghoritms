package main

import "fmt"

type Edge struct {
	to, weight int
}

func main() {
	n := 4
	graph := make([][]Edge, n)

	addEdge := func(a, b, weight int) {
		graph[a] = append(graph[a], Edge{b, weight})
		graph[b] = append(graph[b], Edge{a, weight})
	}

	addEdge(0, 1, 4)
	addEdge(0, 2, 1)
	addEdge(2, 1, 2)
	addEdge(1, 3, 1)

	const INF = int(1e9)
	dist := make([]int, n)
	visited := make([]bool, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	for i := 0; i < n; i++ {
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}
		if u == -1 || dist[u] == INF {
			break
		}

		visited[u] = true
		for _, edge := range graph[u] {
			if dist[u]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[u] + edge.weight
			}
		}
	}

	fmt.Println("Distances from node 0:", dist)
}
