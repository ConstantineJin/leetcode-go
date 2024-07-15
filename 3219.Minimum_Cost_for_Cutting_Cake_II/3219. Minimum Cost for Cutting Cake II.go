package main

import "sort"

type Edge struct {
	cost      int
	isHorizon bool
}

func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	edges := make([]Edge, m+n-2)
	for i, cost := range horizontalCut {
		edges[i] = Edge{cost, true}
	}
	for j, cost := range verticalCut {
		edges[m+j-1] = Edge{cost, false}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].cost > edges[j].cost })
	var totalCost int
	horizPieces, vertPieces := 1, 1
	for _, edge := range edges {
		if edge.isHorizon {
			totalCost += edge.cost * vertPieces
			horizPieces++
		} else {
			totalCost += edge.cost * horizPieces
			vertPieces++
		}
	}
	return int64(totalCost)
}
