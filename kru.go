package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	u, v, weight int
}

func main() {
	edges := []Edge{
		{0, 1, 4}, {0, 2, 1}, {2, 1, 2},
		{1, 3, 1}, {2, 3, 5},
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parent := []int{0, 1, 2, 3}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	totalWeight := 0
	for _, edge := range edges {
		a, b := find(edge.u), find(edge.v)
		if a != b {
			parent[a] = b
			totalWeight += edge.weight
			fmt.Println(edge.u, "--", edge.v, "weight:", edge.weight)
		}
	}

	fmt.Println("Total weight:", totalWeight)
}
