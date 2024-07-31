package main

import "slices"

func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
	slices.SortFunc(points, func(a, b []int) int { return a[0] - b[0] })
	right := -1
	for _, point := range points {
		if point[0] > right {
			right = point[0] + w
			ans++
		}
	}
	return
}
