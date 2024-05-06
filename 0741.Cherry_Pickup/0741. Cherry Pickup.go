package main

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	// f[k][x1][x2]表示两个人分别从(x1,k−x1)和(x2,k−x2)同时出发，到达(N−1,N−1)摘到的樱桃个数之和的最大值
	var dp = make([][][]int, n*2-1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt
			}
		}
	}
	dp[0][0][0] = grid[0][0]
	for k := 1; k < n*2-1; k++ {
		for x1 := max(k-n+1, 0); x1 <= min(k, n-1); x1++ {
			y1 := k - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := x1; x2 <= min(k, n-1); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				res := dp[k-1][x1][x2] // 都往右
				if x1 > 0 {
					res = max(res, dp[k-1][x1-1][x2]) // 往下，往右
				}
				if x2 > 0 {
					res = max(res, dp[k-1][x1][x2-1]) // 往右，往下
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[k-1][x1-1][x2-1]) // 都往下
				}
				res += grid[x1][y1]
				if x2 != x1 { // 避免重复摘同一个樱桃
					res += grid[x2][y2]
				}
				dp[k][x1][x2] = res
			}
		}
	}
	return max(dp[n*2-2][n-1][n-1], 0)
}
