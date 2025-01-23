package main

import "math"

func gridGame(grid [][]int) int64 {
	ans := math.MaxInt
	var left0, left1 int
	for _, v := range grid[0] {
		left0 += v
	}
	for j, v := range grid[0] {
		left0 -= v
		ans = min(ans, max(left0, left1))
		left1 += grid[1][j]
	}
	return int64(ans)
}
