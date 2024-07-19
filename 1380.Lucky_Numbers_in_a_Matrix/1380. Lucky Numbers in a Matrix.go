package main

import "slices"

func luckyNumbers(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	rowMin := make([]int, m)
	for i, row := range matrix {
		rowMin[i] = slices.Min(row)
	}
	for j := 0; j < n; j++ {
		mx := matrix[0][j]
		for i := 1; i < m; i++ {
			mx = max(mx, matrix[i][j])
		}
		for i := 0; i < m; i++ {
			if matrix[i][j] == rowMin[i] && matrix[i][j] == mx {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return
}
