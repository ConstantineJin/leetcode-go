package main

import "slices"

func minOperations(nums []int) int {
	slices.Sort(nums)
	var n = len(nums)
	nums = slices.Compact(nums)
	var ans, left int
	for i, num := range nums {
		for nums[left] < num-n+1 {
			left++
		}
		ans = max(ans, i-left+1)
	}
	return n - ans
}
