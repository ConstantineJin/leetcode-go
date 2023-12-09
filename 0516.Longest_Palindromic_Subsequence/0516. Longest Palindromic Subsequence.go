package main

// 递归
//func longestPalindromeSubseq(s string) int {
//	n := len(s)
//	mem := make([][]int, n) // 记忆化搜索
//	for i := range mem {
//		mem[i] = make([]int, n)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(int, int) int
//	dfs = func(i int, j int) (res int) {
//		if i > j { // 空串
//			return 0
//		} else if i == j { // 只有一个字符，其本身就构成回文子序列
//			return 1
//		}
//		p := &mem[i][j]
//		if *p != -1 { // 已求过当前情况的值
//			return *p
//		}
//		defer func() { *p = res }() // 保存计算结果
//		if s[i] == s[j] {           // 当前首尾字符相同
//			return dfs(i+1, j-1) + 2 // 都选
//		} else { // 当前首尾字符不相同
//			return max(dfs(i, j-1), dfs(i+1, j)) // 动态规划，不选前或者不选后
//		}
//	}
//	return dfs(0, n-1)
//}

// 递推
func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = 1
		pre := 0 // f[i+1][i]
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				pre, f[j] = f[j], pre+2
			} else {
				pre = f[j]
				f[j] = max(f[j], f[j-1])
			}
		}
	}
	return f[n-1]
}
