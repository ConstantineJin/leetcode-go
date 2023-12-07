package main

// 递归
//func minimumDeleteSum(s1 string, s2 string) int {
//	var m, n = len(s1), len(s2)
//	var mem = make([][]int, m+1) // 记忆化搜索
//	for i := range mem {
//		mem[i] = make([]int, n+1)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var sum func(string, int) int // sum函数返回从当前下标开始到字符串末尾所有字符ASCII码之和
//	sum = func(s string, k int) (ans int) {
//		for i := k; i < len(s); i++ {
//			ans += int(s[i])
//		}
//		return
//	}
//	var dfs func(int, int) int
//	dfs = func(i int, j int) (ans int) {
//		var p = &mem[i][j]
//		if *p != -1 { // 已经计算过了
//			return *p
//		}
//		defer func() { *p = ans }() // 保存搜索结果
//		if i == m { // 如果s1遍历完了
//			return sum(s2, j) // s2后面的字符都需要删除
//		}
//		if j == n { // 如果s2遍历完了
//			return sum(s1, i) // s2后面的字符都需要删除
//		}
//		if s1[i] == s2[j] { // 如果当前两个字符相等
//			return dfs(i+1, j+1) // 不需要删除
//		} else {
//			return min(int(s1[i])+dfs(i+1, j), int(s2[j])+dfs(i, j+1)) // 动态规划，删除s1当前所指字符或者删除s2当前所指字符
//		}
//	}
//	return dfs(0, 0)
//}

// 递推
func minimumDeleteSum(s1 string, s2 string) (res int) {
	var m, n = len(s1), len(s2)
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		res += int(s1[i-1])
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}
	for i := range s2 {
		res += int(s2[i])
	}
	return res - dp[m][n]*2
}
