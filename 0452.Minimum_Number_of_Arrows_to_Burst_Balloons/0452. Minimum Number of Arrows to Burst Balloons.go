package main

import "sort"

func findMinArrowShots(points [][]int) (count int) {
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	for len(points) > 0 {
		var boom = 1
		for boom < len(points) && points[boom][0] <= points[0][1] {
			boom++
		}
		points = points[boom:]
		count++
	}
	return
}
