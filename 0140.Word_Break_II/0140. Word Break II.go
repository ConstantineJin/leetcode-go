package main

import "strings"

func wordBreak(s string, wordDict []string) (ans []string) {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	n := len(s)
	var path []string
	var dfs func(prev, i int)
	dfs = func(prev, i int) {
		if i == n {
			if dict[s[prev:n]] {
				path = append(path, s[prev:n])
				ans = append(ans, strings.Join(path, " "))
				path = path[:len(path)-1]
			}
			return
		}
		dfs(prev, i+1)       // 不在当前位置分词
		if dict[s[prev:i]] { // 在当前位置分词
			path = append(path, s[prev:i])
			dfs(i, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1)
	return
}
