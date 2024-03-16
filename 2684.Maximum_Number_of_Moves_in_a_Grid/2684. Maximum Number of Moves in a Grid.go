package main

func maxMoves(grid [][]int) (ans int) {
	var m, n = len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		ans = max(ans, j)
		if j == n-1 {
			return
		}
		for k := max(0, i-1); k < min(m, i+2); k++ {
			if grid[k][j+1] > grid[i][j] {
				dfs(k, j+1)
			}
		}
		grid[i][j] = 0
	}
	for i := range grid {
		dfs(i, 0)
	}
	return
}
