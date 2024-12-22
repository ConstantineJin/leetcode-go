package main

import "slices"

func minSubsequence(nums []int) (ans []int) {
	var s, t int
	for _, num := range nums {
		s += num
	}
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	for _, num := range nums {
		t += num
		ans = append(ans, num)
		if t*2 > s {
			return
		}
	}
	return nums
}
