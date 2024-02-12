package main

func cherryPickup(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var mem = make([][][]int, n)
	for i := range mem {
		mem[i] = make([][]int, n)
		for j := range mem[i] {
			mem[i][j] = make([]int, m)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) (res int) {
		if k == m || i < 0 || j >= n {
			return 0
		}
		var p = &mem[i][j][k]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		switch j - i {
		case 1:
			return max(dfs(i-1, j-1, k+1), dfs(i-1, j, k+1), dfs(i-1, j+1, k+1), dfs(i, j, k+1), dfs(i, j+1, k+1), dfs(i+1, j+1, k+1)) + grid[k][i] + grid[k][j]
		case 2:
			return max(dfs(i-1, j-1, k+1), dfs(i-1, j, k+1), dfs(i-1, j+1, k+1), dfs(i, j, k+1), dfs(i, j+1, k+1), dfs(i+1, j+1, k+1), dfs(i, j-1, k+1), dfs(i+1, j, k+1)) + grid[k][i] + grid[k][j]
		default:
			return max(dfs(i-1, j-1, k+1), dfs(i-1, j, k+1), dfs(i-1, j+1, k+1), dfs(i, j, k+1), dfs(i, j+1, k+1), dfs(i+1, j+1, k+1), dfs(i, j-1, k+1), dfs(i+1, j, k+1), dfs(i+1, j-1, k+1)) + grid[k][i] + grid[k][j]
		}
	}
	return dfs(0, n-1, 0)
}
