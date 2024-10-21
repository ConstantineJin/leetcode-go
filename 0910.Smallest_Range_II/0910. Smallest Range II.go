package main

import "sort"

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 1; i < n; i++ {
		ans = min(ans, max(nums[i-1]+k, nums[n-1]-k)-min(nums[0]+k, nums[i]-k))
	}
	return ans
}
