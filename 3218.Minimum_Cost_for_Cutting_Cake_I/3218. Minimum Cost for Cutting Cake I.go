package main

import "sort"

type Edge struct {
	cost      int
	isHorizon bool
}

func minimumCost(m, n int, horizontalCut, verticalCut []int) (ans int) {
	edges := make([]Edge, m+n-2)
	for i, cost := range horizontalCut {
		edges[i] = Edge{cost, true}
	}
	for j, cost := range verticalCut {
		edges[m+j-1] = Edge{cost, false}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].cost > edges[j].cost })
	horizPieces, vertPieces := 1, 1
	for _, edge := range edges {
		if edge.isHorizon {
			ans += edge.cost * vertPieces
			horizPieces++
		} else {
			ans += edge.cost * horizPieces
			vertPieces++
		}
	}
	return
}
