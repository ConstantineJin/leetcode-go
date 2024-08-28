package main

import "math"

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	mem := make([]int, n)
	var dfs func(i int) int // dfs(i) 表示 s[i:] 最少能分割出多少个平衡子串
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		if mem[i] > 0 {
			return mem[i]
		}
		res := math.MaxInt
		var cnt [26]int
		var k, maxCnt int // 子串中有 k 种字母，字母的频数最大值为 maxCnt
		for j := i; j < n; j++ {
			c := s[j] - 'a'
			if cnt[c] == 0 {
				k++
			}
			cnt[c]++
			maxCnt = max(maxCnt, cnt[c])
			if j-i+1 == k*maxCnt { // 子串是平衡的充要条件是子串长度等于子串中的字母频数最大值乘以子串中不同字母的个数
				res = min(res, dfs(j+1)+1)
			}
		}
		mem[i] = res
		return res
	}
	return dfs(0)
}
