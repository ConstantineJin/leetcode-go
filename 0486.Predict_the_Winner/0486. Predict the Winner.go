package main

func predictTheWinner(nums []int) bool {
	var n = len(nums)
	dp := make([][]int, n) // dp[i][j]表示当数组剩下的部分为下标i到下标j时，即在下标范围[i,j]中，当前玩家与另一个玩家的分数之差的最大值，注意当前玩家不一定是先手
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}
