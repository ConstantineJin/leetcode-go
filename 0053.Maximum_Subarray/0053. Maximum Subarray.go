package main

import "math"

func maxSubArray(nums []int) int {
	n, left, right, ans := len(nums), 0, 0, math.MinInt
	for left < n {
		if nums[left] <= 0 {
			left++
			continue
		}
		sum := nums[left]
		ans = max(ans, sum)
		for right = left + 1; right < n && sum+nums[right] >= 0; right++ { // 如果当前子数组和已经为负，则不可能成为和最大的子数组
			sum += nums[right]
			ans = max(ans, sum)
		}
		left = right + 1
	}
	if ans == math.MinInt { // nums所有元素全是负数
		for _, num := range nums {
			ans = max(ans, num) // 找其中最大的负数作为结果返回
		}
	}
	return ans
}
