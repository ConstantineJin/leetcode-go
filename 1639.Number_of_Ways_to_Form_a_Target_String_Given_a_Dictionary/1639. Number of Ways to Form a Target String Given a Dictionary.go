package main

const mod int = 1e9 + 7

func numWays(words []string, target string) int {
	m, n := len(words[0]), len(target)
	dp := make([][]int, m+1)    // dp[i][j] 表示 words 中前 i 个字符，拼成 target 前 j 个字符的方案数
	cnt := make([][26]int, m+1) // cnt[i][c] 表示在 words 中第 i 个字符是字符 c 的单词个数
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for _, word := range words {
			cnt[i][int(word[i-1]-'a')]++
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]*cnt[i][int(target[j-1]-'a')]) % mod
		}
	}
	return dp[m][n]
}
