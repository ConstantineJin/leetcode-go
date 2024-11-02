package main

func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	canReach := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]int, n)
		canReach[i] = make([]bool, n)
		canReach[i][0] = true
	}
	for j := 1; j < n; j++ {
		for i := range m {
			if i > 0 && canReach[i-1][j-1] && grid[i][j] > grid[i-1][j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
				canReach[i][j] = true
			}
			if canReach[i][j-1] && grid[i][j] > grid[i][j-1] {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+1)
				canReach[i][j] = true
			}
			if i < m-1 && canReach[i+1][j-1] && grid[i][j] > grid[i+1][j-1] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+1)
				canReach[i][j] = true
			}
			if canReach[i][j] {
				ans = max(ans, dp[i][j])
			}
		}
	}
	return
}
