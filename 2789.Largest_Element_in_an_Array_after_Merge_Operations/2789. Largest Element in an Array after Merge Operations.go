package main

func maxArrayValue(nums []int) int64 {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] <= nums[i] {
			nums[i-1] += nums[i] // 把合并值向左传
		}
	}
	return int64(nums[0])
}
