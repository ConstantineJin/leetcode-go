package main

import "math"

func maxRemovals(source, pattern string, targetIndices []int) int {
	n, m := len(source), len(pattern)
	isTarget := make([]int, n)
	for _, index := range targetIndices {
		isTarget[index] = 1
	}
	dp := make([][]int, n+1) // dp[i][j] 表示使 pattern[0:j] 是 source[0:i] 的子序列，最多可以进行多少次删除操作
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	dp[0][0] = 0
	for i := range source {
		isDel := isTarget[i]
		dp[i+1][0] = dp[i][0] + isDel
		for j := range min(i+1, m) {
			res := dp[i][j+1] + isDel    // 不选 source[i]。如果 i 在 targetIndices 中需要将删除次数 +1
			if source[i] == pattern[j] { // 如果 source[i] 与 pattern[j] 相同，可以都选，转移到 dp[i][j]
				res = max(res, dp[i][j])
			}
			dp[i+1][j+1] = res
		}
	}
	return dp[n][m]
}
