package main

func transpose(matrix [][]int) [][]int {
	var m, n = len(matrix), len(matrix[0])
	var ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[j][i] = matrix[i][j]
		}
	}
	return ans
}
