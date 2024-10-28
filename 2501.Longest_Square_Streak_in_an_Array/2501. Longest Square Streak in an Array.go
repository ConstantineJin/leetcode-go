package main

import (
	"slices"
	"sort"
)

func longestSquareStreak(nums []int) (ans int) {
	sort.Ints(nums)
	nums = slices.Compact(nums)
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	for _, num := range nums {
		var cnt int
		for set[num] {
			num *= num // 题目保证 num >= 2，故无需担心 0 和 1 造成死循环
			cnt++
		}
		ans = max(ans, cnt)
	}
	if ans < 2 {
		return -1
	}
	return ans
}
