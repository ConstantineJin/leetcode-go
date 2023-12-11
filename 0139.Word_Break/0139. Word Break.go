package main

import "strings"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	mem := make([]int, n+1) // 记忆化搜索，0表示未计算过从当前下标开始到原字符串末尾的子串是否完成单词拆分
	var dfs func(i int) bool
	dfs = func(i int) (flag bool) {
		if i == n { // 可以利用字典中的单词拼接到原串末尾
			return true
		}
		if mem[i] == 1 { // 已经计算过，从当前下标开始到原字符串末尾的子串可以完成单词拆分
			return true
		} else if mem[i] == -1 { // 已经计算过，从当前下标开始到原字符串末尾的子串不能完成单词拆分
			return false
		}
		for _, word := range wordDict {
			if strings.HasPrefix(s[i:], word) { // 只要字典中的某个单词是当前剩余子串的开头
				flag = flag || dfs(i+len(word)) // dfs
			}
		}
		if flag { // 保存搜索结果
			mem[i] = 1
		} else {
			mem[i] = -1
		}
		return flag
	}
	return dfs(0)
}
