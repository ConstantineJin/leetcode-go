package main

func restoreMatrix(rowSum, colSum []int) [][]int {
	matrix := make([][]int, len(rowSum))
	for i, rs := range rowSum {
		matrix[i] = make([]int, len(colSum))
		for j, cs := range colSum {
			matrix[i][j] = min(rs, cs)
			rs -= matrix[i][j]
			colSum[j] -= matrix[i][j]
		}
	}
	return matrix
}
