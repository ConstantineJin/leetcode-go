package main

// 递归
//func longestCommonSubsequence(text1 string, text2 string) int {
//	var m, n = len(text1), len(text2)
//	var mem = make([][]int, m)
//	for i := range mem {
//		mem[i] = make([]int, n)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(i, j int) int
//	dfs = func(i, j int) (res int) {
//		if i == m || j == n {
//			return 0
//		}
//		var p = &mem[i][j]
//		if *p != -1 {
//			return *p
//		}
//		defer func() { *p = res }()
//		if text1[i] == text2[j] {
//			return dfs(i+1, j+1) + 1
//		}
//		return max(dfs(i+1, j), dfs(i, j+1))
//	}
//	return dfs(0, 0)
//}

// 递推
//func longestCommonSubsequence(text1 string, text2 string) int {
//	var m, n = len(text1), len(text2)
//	var f = make([][]int, m+1)
//	for i := range f {
//		f[i] = make([]int, n+1)
//	}
//	for i, x := range text1 {
//		for j, y := range text2 {
//			if x == y {
//				f[i+1][j+1] = f[i][j] + 1
//			} else {
//				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
//			}
//		}
//	}
//	return f[m][n]
//}

// 递推空间优化
func longestCommonSubsequence(text1, text2 string) int {
	var n = len(text2)
	var f = make([]int, n+1)
	for _, x := range text1 {
		var pre int
		for j, y := range text2 {
			if x == y {
				f[j+1], pre = pre+1, f[j+1]
			} else {
				pre = f[j+1]
				f[j+1] = max(f[j+1], f[j])
			}
		}
	}
	return f[n]
}
