package main

func satisfiesConditions(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if i < m-1 && v != grid[i+1][j] || j < n-1 && v == row[j+1] {
				return false
			}
		}
	}
	return true
}
