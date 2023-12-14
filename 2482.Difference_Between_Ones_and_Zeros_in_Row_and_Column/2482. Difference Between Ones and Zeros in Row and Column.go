package main

func onesMinusZeros(grid [][]int) [][]int {
	var m, n = len(grid), len(grid[0])
	var onesRow, onesCol = make([]int, m), make([]int, n)
	for i, r := range grid {
		for j, v := range r {
			onesRow[i] += v*2 - 1 // 把0当作-1
			onesCol[j] += v*2 - 1
		}
	}
	for i, x := range onesRow {
		for j, y := range onesCol {
			grid[i][j] = x + y
		}
	}
	return grid
}
