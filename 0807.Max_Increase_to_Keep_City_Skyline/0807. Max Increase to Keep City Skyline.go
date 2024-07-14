package main

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	n := len(grid)
	rowMax, colMax := make([]int, n), make([]int, n)
	for i, row := range grid {
		for j, v := range row {
			rowMax[i], colMax[j] = max(rowMax[i], v), max(colMax[j], v)
		}
	}
	for i, row := range grid {
		for j, v := range row {
			ans += min(rowMax[i], colMax[j]) - v
		}
	}
	return
}
