package main

func kInversePairs(n int, k int) int {
	prevDP := make([]int, k+1)
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 2; i <= n; i++ {
		prevDP, dp = dp, prevDP
		var sum int
		for j := range dp {
			sum += prevDP[j]
			if j >= i {
				sum -= prevDP[j-i]
			}
			dp[j] = sum % (1e9 + 7)
		}
	}
	return dp[k]
}
