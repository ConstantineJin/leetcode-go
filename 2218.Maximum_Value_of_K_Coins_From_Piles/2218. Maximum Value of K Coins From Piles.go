package main

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	dp := make([][]int, n+1) // dp[i+1][j] 表示从 piles[0] 到 piles[i] 中选取体积之和至多 j 个硬币的面值最大和
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i, pile := range piles {
		for j := range k + 1 {
			dp[i+1][j] = dp[i][j] // 不选这一组
			var v int
			for w := range min(j, len(pile)) {
				v += pile[w]
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-w-1]+v)
			}
		}
	}
	return dp[n][k]
}
