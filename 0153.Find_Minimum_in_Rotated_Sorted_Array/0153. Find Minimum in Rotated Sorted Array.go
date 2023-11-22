package main

func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间二分查找
	for left+1 < right {
		middle := (left + right) / 2
		if nums[middle] < nums[len(nums)-1] { // 中间的元素比右端的小，说明最小的元素在中间元素的左边
			right = middle
		} else { // 否则最小的元素在中间元素的右边
			left = middle
		}
	}
	return nums[right]
}
