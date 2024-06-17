package main

import "sort"

// 如果存在特殊序列，那么该序列一定是 strs 中的某个完整字符串
func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool { return len(strs[i]) > len(strs[j]) })
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubsequence(s, t) {
				continue next
			}
		}
		return len(strs[i])
	}
	return -1
}

func isSubsequence(s, t string) bool {
	m, n := len(s), len(t)
	if m > n {
		return false
	}
	var i int
	for j := 0; i < m && j < n; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == m
}
