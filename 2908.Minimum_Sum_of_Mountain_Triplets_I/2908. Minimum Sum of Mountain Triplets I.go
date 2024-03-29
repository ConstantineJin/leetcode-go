package main

import "math"

func minimumSum(nums []int) int {
	var n = len(nums)
	var suf = make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}
	var ans, pre = math.MaxInt, nums[0]
	for i := 1; i < n-1; i++ {
		if nums[i] > pre && nums[i] > suf[i+1] {
			ans = min(ans, pre+nums[i]+suf[i+1])
		}
		pre = min(pre, nums[i])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
