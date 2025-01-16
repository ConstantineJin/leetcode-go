package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, num := range nums {
		rightOr |= num
		for left <= right && nums[left]|rightOr >= k {
			ans = min(ans, right-left+1)
			left++
			if bottom < left {
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
