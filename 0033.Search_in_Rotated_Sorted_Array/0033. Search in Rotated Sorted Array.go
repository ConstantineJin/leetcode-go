package main

// 两次二分
// findMin 在nums中利用二分法寻找最小值下标（见153题）
//func findMin(nums []int) int {
//	left, right := -1, len(nums)-1 // 开区间二分查找
//	for left+1 < right {
//		middle := (left + right) / 2
//		if nums[middle] < nums[len(nums)-1] { // 中间的元素比右端的小，说明最小的元素在中间元素的左边
//			right = middle
//		} else { // 否则最小的元素在中间元素的右边
//			left = middle
//		}
//	}
//	return right
//}
//
//func lowerBound(nums []int, left, right, target int) int {
//	r0 := right
//	for left+1 < right { // 开区间不为空
//		middle := left + (right-left)/2
//		if nums[middle] < target {
//			left = middle // 范围缩小到 (middle, right)
//		} else {
//			right = middle // 范围缩小到 (left, middle)
//		}
//	}
//	if right == r0 || nums[right] != target {
//		return -1
//	}
//	return right
//}
//
//func search(nums []int, target int) int {
//	i := findMin(nums)
//	if target > nums[len(nums)-1] {
//		return lowerBound(nums, -1, i, target) // 左段
//	}
//	return lowerBound(nums, i-1, len(nums), target) // 右段
//}

// 一次二分
func search(nums []int, target int) int {
	isBlue := func(i int) bool {
		end := nums[len(nums)-1]
		if nums[i] > end {
			return target > end && nums[i] >= target
		} else {
			return target > end || nums[i] >= target
		}
	}
	left, right := -1, len(nums) // 开区间
	for left+1 < right {
		middle := (left + right) / 2
		if isBlue(middle) {
			right = middle
		} else {
			left = middle
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}
