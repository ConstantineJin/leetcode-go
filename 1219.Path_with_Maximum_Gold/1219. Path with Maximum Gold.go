package main

var dirs = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func getMaximumGold(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j, gold int)
	dfs = func(i, j, gold int) {
		gold += grid[i][j]
		ans = max(ans, gold)
		temp := grid[i][j]
		grid[i][j] = 0
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] > 0 {
				dfs(x, y, gold)
			}
		}
		grid[i][j] = temp
	}
	for i, row := range grid {
		for j, v := range row {
			if v > 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}
