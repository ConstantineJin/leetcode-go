package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}
	sum[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0] // 初始化第一列，只能经由上方的格子抵达
	}
	for j := 1; j < n; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j] // 初始化第一行，只能经由左侧的格子抵达
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = grid[i][j] + min(sum[i-1][j], sum[i][j-1]) // 到达每个格子的路径只有两种：从上往下或从左往右，因此选择每个格子上方和左侧的格子中较小者
		}
	}
	return sum[m-1][n-1]
}
