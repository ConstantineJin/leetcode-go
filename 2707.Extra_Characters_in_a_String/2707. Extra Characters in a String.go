package main

import (
	"strings"
)

// 动态规划，递归
//func minExtraChar(s string, dictionary []string) int {
//	var n = len(s)
//	var mem = make([]int, n)
//	for i := range mem {
//		mem[i] = -1
//	}
//	var dfs func(i int) int // dfs(i)表示前i+1个字符的子问题
//	dfs = func(i int) (res int) {
//		if i < 0 {
//			return
//		}
//		var p = &mem[i]
//		if *p != -1 {
//			return *p
//		}
//		res = dfs(i-1) + 1                // 不选当前字符
//		for _, word := range dictionary { // 枚举dictionary中的单词
//			if strings.HasSuffix(s[:i+1], word) { // 如果当前子串以dictionary中的单词作为结尾
//				res = min(res, dfs(i-len(word))) // dfs并更新结果
//			}
//		}
//		*p = res
//		return res
//	}
//	return dfs(n - 1)
//}

// 动态规划，递推
func minExtraChar(s string, dictionary []string) int {
	var n = len(s)
	var f = make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1 // 不选当前字符
		for _, word := range dictionary {
			if strings.HasSuffix(s[:i+1], word) {
				f[i+1] = min(f[i+1], f[i-len(word)+1])
			}
		}
	}
	return f[n]
}
