package main

import (
	"fmt"
	"sort"
)

// Edge represents a single graph edge.
type Edge struct {
	From   int
	To     int
	Weight int
}

// UnionFind (Disjoint Set) tracks which nodes already belong to the
// same group, so cycles can be detected while building the MST.
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		uf.parent[i] = i
	}
	return uf
}

// Find returns the root of a node and applies path compression.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges two groups. It returns false when both nodes are already
// in the same group, which means the edge would create a cycle.
func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)
	if rootA == rootB {
		return false
	}

	if uf.rank[rootA] < uf.rank[rootB] {
		rootA, rootB = rootB, rootA
	}
	uf.parent[rootB] = rootA
	if uf.rank[rootA] == uf.rank[rootB] {
		uf.rank[rootA]++
	}
	return true
}

// Kruskal returns the Minimum Spanning Tree and its total weight.
func Kruskal(edges []Edge, numNodes int) ([]Edge, int) {
	// 1. Sort all edges by weight, cheapest first.
	sorted := make([]Edge, len(edges))
	copy(sorted, edges)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight < sorted[j].Weight
	})

	uf := NewUnionFind(numNodes)
	mst := []Edge{}
	total := 0

	// 2. Take every edge that does not create a cycle.
	for _, e := range sorted {
		if uf.Union(e.From, e.To) {
			mst = append(mst, e)
			total += e.Weight
		}
	}

	return mst, total
}

func main() {
	edges := []Edge{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 3},
		{2, 4, 4},
		{3, 4, 5},
		{3, 5, 6},
		{4, 5, 7},
	}
	numNodes := 5

	mst, total := Kruskal(edges, numNodes)

	fmt.Println("MST edges:")
	for _, e := range mst {
		fmt.Printf("  %d - %d  weight %d\n", e.From, e.To, e.Weight)
	}

	fmt.Println("Total weight:", total)

	if len(mst) == numNodes-1 {
		fmt.Println("Graph is connected, MST found.")
	} else {
		fmt.Println("Graph is not connected.")
	}
}
