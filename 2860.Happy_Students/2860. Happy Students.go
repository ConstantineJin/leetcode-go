package main

import "sort"

func countWays(nums []int) (ans int) {
	sort.Ints(nums)
	if nums[0] > 0 { // 如果最小值为正，那么一定可以都不选
		ans = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < i && i < nums[i] {
			ans++
		}
	}
	return ans + 1 // 题目数据范围保证 nums[i]<len(nums)，因此一定可以都选
}
