package main

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			for m := 0; m < i; m++ {
				if nums[m] == nums[i] {
					dp[i][j] = max(dp[i][j], dp[m][j]+1)
				} else if j > 0 {
					dp[i][j] = max(dp[i][j], dp[m][j-1]+1)
				}
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen
}
