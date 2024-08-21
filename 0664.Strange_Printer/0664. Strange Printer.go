package main

import "math"

func strangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n) // f[i][j] 表示打印完成区间 [i,j] 的最少操作数
	for i := range f {
		f[i] = make([]int, n)
		f[i][i] = 1 // 一个字符只需要一次打印
	}
	for i := n - 1; i >= 0; i-- { // 倒序枚举 i，因为计算 f[i][j] 时需要用到 f[k+1][j]，其中 k>=i
		for j := i + 1; j < n; j++ { // 正序枚举 j，因为计算 f[i][j] 时需要用到 f[i][k]，其中 k<=j
			if s[i] == s[j] { // 如果当前子串首尾相同，那么不需要消耗额外的打印次数
				f[i][j] = f[i][j-1]
			} else {
				f[i][j] = math.MaxInt
				for k := i; k < j; k++ {
					f[i][j] = min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
	}
	return f[0][n-1]
}
