package main

func countServers(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	rows, cols := make([]int, m), make([]int, n)
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 && (rows[i] > 1 || cols[j] > 1) {
				ans++
			}
		}
	}
	return
}
