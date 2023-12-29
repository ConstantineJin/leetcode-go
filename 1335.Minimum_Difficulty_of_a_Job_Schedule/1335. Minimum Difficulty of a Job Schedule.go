package main

import "math"

// DP，递归
//func minDifficulty(jobDifficulty []int, d int) int {
//	var n = len(jobDifficulty)
//	if n < d {
//		return -1
//	}
//	var mem = make([][]int, n)
//	for i := range mem {
//		mem[i] = make([]int, d)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(int, int) int
//	dfs = func(i, j int) (res int) { // 前i+1项工作在前j+1天做完的最小难度之和
//		var p = &mem[i][j]
//		if *p != -1 {
//			return *p
//		}
//		defer func() { *p = res }()
//		if j == 0 { // 如果只剩一天
//			return slices.Max(jobDifficulty[:i+1]) // 需要把前面所剩的工作都做完
//		}
//		var today = jobDifficulty[i] // 每天至少做一项工作
//		res = jobDifficulty[i] + dfs(i-1, j-1)
//		for k := i; k >= j; k-- {
//			today = max(today, jobDifficulty[k]) // 今天的工作难度
//			res = min(res, dfs(k-1, j-1)+today)  // 今天的和之前的工作难度之和的最小值
//		}
//		return
//	}
//	return dfs(n-1, d-1)
//}

// DP，递推
//func minDifficulty(jobDifficulty []int, d int) int {
//	var n = len(jobDifficulty)
//	if n < d {
//		return -1
//	}
//	var f = make([][]int, d)
//	f[0] = make([]int, n)
//	f[0][0] = jobDifficulty[0]
//	for i := 1; i < n; i++ {
//		f[0][i] = max(f[0][i-1], jobDifficulty[i])
//	}
//	for i := 1; i < d; i++ {
//		f[i] = make([]int, n)
//		for j := n - 1; j >= i; j-- {
//			f[i][j] = math.MaxInt
//			var mx int
//			for k := j; k >= i; k-- {
//				mx = max(mx, jobDifficulty[k])
//				f[i][j] = min(f[i][j], f[i-1][k-1]+mx)
//			}
//		}
//	}
//	return f[d-1][n-1]
//}

// DP，递推空间优化
func minDifficulty(jobDifficulty []int, d int) int {
	var n = len(jobDifficulty)
	if n < d {
		return -1
	}
	var f = make([]int, n)
	f[0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1], jobDifficulty[i])
	}
	for i := 1; i < d; i++ {
		for j := n - 1; j >= i; j-- {
			f[j] = math.MaxInt
			var mx int
			for k := j; k >= i; k-- {
				mx = max(mx, jobDifficulty[k]) // 从 a[k] 到 a[j] 的最大值
				f[j] = min(f[j], f[k-1]+mx)
			}
		}
	}
	return f[n-1]
}
