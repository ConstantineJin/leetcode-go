package main

import "math"

func maxScore(a, b []int) int64 {
	n := len(b)
	dp := make([][5]int64, n+1)
	for j := 1; j < 5; j++ {
		dp[0][j] = math.MinInt64 >> 1
	}
	for i, y := range b {
		for j, x := range a {
			dp[i+1][j+1] = max(dp[i][j+1], dp[i][j]+int64(x*y))
		}
	}
	return dp[n][4]
}
