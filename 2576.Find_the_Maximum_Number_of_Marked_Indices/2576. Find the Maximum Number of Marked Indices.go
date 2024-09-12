package main

import "sort"

func maxNumOfMarkedIndices(nums []int) (ans int) {
	sort.Ints(nums)
	for _, num := range nums[(len(nums)+1)/2:] { // 最多匹配 n/2 对，使用双指针匹配
		if nums[ans]*2 <= num {
			ans++
		}
	}
	return ans * 2
}
