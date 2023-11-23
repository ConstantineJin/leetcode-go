package main

// 顺序遍历寻找最大值，最大值一定是峰值，时间复杂度O(n)
//func findPeakElement(nums []int) (ans int) {
//	var m = nums[0]
//	for i, num := range nums {
//		if num > m {
//			m, ans = num, i
//		}
//	}
//	return
//}

// 二分查找，开区间
func findPeakElement(nums []int) int {
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
