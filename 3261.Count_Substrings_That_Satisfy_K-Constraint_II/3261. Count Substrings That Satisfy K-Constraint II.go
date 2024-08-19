package main

import "sort"

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n) // 以 i 为右端点的合法子串，其左端点最小是 left[i]。显然 left 具有单调性
	prefixSum := make([]int, n+1)
	var cnt [2]int
	var l int
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		prefixSum[i+1] = prefixSum[i] + i - l + 1 // 计算 i-left[i]+1 的前缀和
	}
	ans := make([]int64, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		j := l + sort.SearchInts(left[l:r+1], l)                        // 在 [l,r] 中二分查找 left 中的第一个满足 left[j]>=l 的下标 j
		ans[i] = int64(prefixSum[r+1] - prefixSum[j] + (j-l+1)*(j-l)/2) // [l,j-1] 的所有子串都是合法的。右端点在 [j,r] 内的子串，可以累加下标在 [j,r] 内的所有 i−left[i]+1 的和
	}
	return ans
}
