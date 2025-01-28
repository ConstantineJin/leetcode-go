package main

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y]
		grid[x][y] = 0
		for _, dir := range dirs {
			sum += dfs(x+dir[0], y+dir[1])
		}
		return sum
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}
