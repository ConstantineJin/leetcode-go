package main

func search(nums []int, target int) bool {
	left, right := -1, len(nums)-1
	check := func(i int) bool {
		num := nums[i]
		if num > nums[right] {
			return target > nums[right] && num >= target
		}
		return target > nums[right] || num >= target
	}
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] == nums[right] {
			right--
		} else if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right] == target
}
