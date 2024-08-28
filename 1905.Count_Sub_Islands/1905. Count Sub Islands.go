package main

var dirs = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func countSubIslands(grid1, grid2 [][]int) (ans int) {
	m, n := len(grid1), len(grid1[0])
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if grid2[i][j] == 0 {
			return true
		}
		grid2[i][j] = 0 // 标记当前格子已访问
		res := grid1[i][j] == 1
		for _, dir := range dirs {
			if x, y := i+dir[0], j+dir[1]; 0 <= x && x < m && 0 <= y && y < n && !dfs(x, y) {
				res = false
			}
		}
		return res
	}
	for i, row := range grid2 {
		for j, cell := range row {
			if cell == 1 {
				if dfs(i, j) {
					ans++
				}
			}
		}
	}
	return
}
