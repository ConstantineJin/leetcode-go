package main

import "sort"

func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums) // 排序
	for i := 1; i < len(nums); i++ {
		ans += sort.SearchInts(nums[0:i], target-nums[i]) // 二分查找
	}
	return
}
