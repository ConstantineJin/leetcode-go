package main

import "sort"

func minDifference(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < 4; i++ {
		ans = min(ans, nums[n+i-4]-nums[i])
	}
	return ans
}
