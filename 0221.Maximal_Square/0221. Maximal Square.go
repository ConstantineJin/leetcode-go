package main

func maximalSquare(matrix [][]byte) (ans int) {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m) // dp二维数组维护以当前节点为右下角的全1正方形的边长
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				ans = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1 // 递推关系：左上方、上方和左侧三格的最小值+1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}
