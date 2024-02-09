package main

// DFS
func numIslands(grid [][]byte) (ans int) {
	var m, n = len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i int, j int) {
		grid[i][j] = 0
		if i > 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if j > 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if i < m-1 && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j < n-1 && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
