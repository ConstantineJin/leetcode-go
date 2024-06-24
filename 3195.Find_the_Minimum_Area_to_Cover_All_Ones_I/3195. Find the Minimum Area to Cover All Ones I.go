package main

func minimumArea(grid [][]int) int {
	minRow, minCol := len(grid)-1, len(grid[0])-1
	var maxRow, maxCol int
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				minRow = min(minRow, i)
				minCol = min(minCol, j)
				maxRow = max(maxRow, i)
				maxCol = max(maxCol, j)
			}
		}
	}
	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
