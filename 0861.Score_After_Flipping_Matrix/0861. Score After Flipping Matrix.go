package main

import "math"

func matrixScore(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		if row[0] == 0 {
			for j := range row {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	for j := range grid[0] {
		var oneCnt int
		for i := range grid {
			if grid[i][j] == 1 {
				oneCnt++
			}
		}
		if oneCnt <= m/2 {
			for i := range grid {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	for _, row := range grid {
		for j, v := range row {
			ans += v * int(math.Pow(2, float64(n-j-1)))
		}
	}
	return
}
