package main

import "cmp"

func countMatchingSubarrays(nums, pattern []int) (ans int) {
	n, m := len(nums), len(pattern)
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	for i, j := 1, 0; i < n; i++ {
		v := cmp.Compare(nums[i], nums[i-1])
		for j > 0 && pattern[j] != v {
			j = next[j-1]
		}
		if pattern[j] == v {
			j++
		}
		if j == m {
			ans++
			j = next[j-1]
		}
	}
	return
}
