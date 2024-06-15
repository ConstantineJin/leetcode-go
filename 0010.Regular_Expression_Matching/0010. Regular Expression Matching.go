package main

import (
	"regexp"
	"strings"
)

// 动态规划，递归
//func isMatch(s, p string) bool {
//	m, n := len(s), len(p)
//	var dfs func(i, j int) bool
//	dfs = func(i, j int) bool {
//		if i == m && j == n { // s 和 p 都已经遍历完成，匹配成功
//			return true
//		}
//		if i < m && j == n { // p 已经遍历完但 s 还有剩余，匹配失败
//			return false
//		}
//		if j < n-1 && p[j+1] == '*' { // 下一个字符是'*'
//			if dfs(i, j+2) { // '*' 表示0个字符
//				return true
//			}
//			for i < m && (p[j] == '.' || s[i] == p[j]) { // '*'表示多个字符
//				if dfs(i+1, j+2) {
//					return true
//				}
//				i++
//			}
//		} else { // 下一个字符不是'*'，只需考虑 s[i] 和 p[j] 是否匹配
//			if i < m && (p[j] == '.' || s[i] == p[j]) {
//				return dfs(i+1, j+1)
//			}
//		}
//		return false
//	}
//	return dfs(0, 0)
//}

// 动态规划，递推
//func isMatch(s string, p string) bool {
//	m, n := len(s), len(p)
//	matches := func(i, j int) bool {
//		if i == 0 {
//			return false
//		}
//		if p[j-1] == '.' {
//			return true
//		}
//		return s[i-1] == p[j-1]
//	}
//	f := make([][]bool, m + 1)
//	for i := 0; i < len(f); i++ {
//		f[i] = make([]bool, n + 1)
//	}
//	f[0][0] = true
//	for i := 0; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			if p[j-1] == '*' {
//				f[i][j] = f[i][j] || f[i][j-2]
//				if matches(i, j - 1) {
//					f[i][j] = f[i][j] || f[i-1][j]
//				}
//			} else if matches(i, j) {
//				f[i][j] = f[i][j] || f[i-1][j-1]
//			}
//		}
//	}
//	return f[m][n]
//}

// 标准库
func isMatch(s, p string) bool {
	a := strings.FieldsFunc("^"+p+"$", func(r rune) bool { return r == '*' })
	p = strings.Join(a, "*")
	return regexp.MustCompile(p).MatchString(s)
}
