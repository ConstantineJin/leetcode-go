package main

import (
	"math"
	"slices"
)

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n) // 将两端也视作切开点
	slices.Sort(cuts)
	m := len(cuts)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := m - 3; i >= 0; i-- {
		for j := i + 2; j < m; j++ {
			res := math.MaxInt
			for k := i + 1; k < j; k++ { // 枚举从哪个点切开
				res = min(res, dp[i][k]+dp[k][j])
			}
			dp[i][j] = res + cuts[j] - cuts[i] // 切这一刀的成本是 cuts[j] - cuts[i]
		}
	}
	return dp[0][m-1]
}
