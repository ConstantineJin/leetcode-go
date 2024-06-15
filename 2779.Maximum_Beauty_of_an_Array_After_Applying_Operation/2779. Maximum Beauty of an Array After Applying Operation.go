package main

import "sort"

func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums) // 选择的子序列中各元素相等，因此元素顺序不影响答案，可以先排序
	var left int
	for right, num := range nums { // 滑动窗口
		for num-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
