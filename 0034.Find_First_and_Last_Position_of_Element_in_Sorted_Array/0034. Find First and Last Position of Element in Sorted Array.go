package main

// 二分查找寻找nums中第一个等于target的下标
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 闭区间
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else { // 即使相等右边界也要左移
			right = middle - 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target { // 如果target已经超过了数组最大值或者不在数组中
		return []int{-1, -1} // 直接返回{-1, -1}
	}
	return []int{start, lowerBound(nums, target+1) - 1} // 右边界是target+1在nums中第一个可能出现的位置左移一位
}
