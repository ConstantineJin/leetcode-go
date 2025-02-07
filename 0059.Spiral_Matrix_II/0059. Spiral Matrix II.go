package main

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, n)
	}
	var row, col, idx int
	for i := 1; i <= n*n; i++ {
		ans[row][col] = i
		dir := dirs[idx]
		if r, c := row+dir[0], col+dir[1]; r < 0 || r >= n || c < 0 || c >= n || ans[r][c] > 0 {
			idx = (idx + 1) % 4
			dir = dirs[idx]
		}
		row += dir[0]
		col += dir[1]
	}
	return ans
}
