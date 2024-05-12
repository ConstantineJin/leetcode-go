package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := range ans {
		ans[i] = make([]int, n-2)
	}
	dirs := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			mx := grid[i][j]
			for _, dir := range dirs {
				mx = max(mx, grid[i+dir[0]][j+dir[1]])
			}
			ans[i-1][j-1] = mx
		}
	}
	return ans
}
