package main

func countSquares(matrix [][]int) (ans int) {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = matrix[i]
		for j := range dp[i] {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	for _, row := range dp {
		for _, v := range row {
			ans += v
		}
	}
	return
}
