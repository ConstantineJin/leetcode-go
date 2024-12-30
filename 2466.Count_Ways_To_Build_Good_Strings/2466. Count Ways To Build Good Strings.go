package main

const mod int = 1e9 + 7

func countGoodStrings(low, high, zero, one int) (ans int) {
	dp := make([]int, high+1) // dp[i] 表示长度为 i 的好字符串数量
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = dp[i-zero]
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}
	return
}
