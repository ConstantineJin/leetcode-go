package main

func modifiedMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	mx := make([]int, n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			mx[j] = max(mx[j], matrix[i][j])
		}
	}
	ans := make([][]int, m)
	for i, row := range matrix {
		ans[i] = make([]int, n)
		for j, v := range row {
			if v == -1 {
				ans[i][j] = mx[j]
			} else {
				ans[i][j] = matrix[i][j]
			}
		}
	}
	return ans
}
