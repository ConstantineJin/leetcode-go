package main

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	var sum, i int
	for _, num := range nums {
		sum += num
	}
	for i = len(nums) - 1; i >= 0 && sum-nums[i] <= nums[i]; i-- {
		sum -= nums[i]
	}
	if sum == 0 {
		return -1
	}
	return int64(sum)
}
