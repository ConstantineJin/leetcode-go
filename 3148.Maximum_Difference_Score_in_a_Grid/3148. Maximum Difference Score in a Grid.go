package main

import "math"

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := math.MinInt
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for j := range f[0] {
		f[0][j] = math.MaxInt
	}
	for i, row := range grid {
		f[i+1] = make([]int, n+1)
		f[i+1][0] = math.MaxInt
		for j, v := range row {
			mn := min(f[i][j+1], f[i+1][j])
			ans = max(ans, v-mn)
			f[i+1][j+1] = min(mn, v)
		}
	}
	return ans
}
