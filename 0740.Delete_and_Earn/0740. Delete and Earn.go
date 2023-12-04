package main

func deleteAndEarn(nums []int) int {
	maxVal := 0 // nums中最大的元素
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1) // 记录每个相同元素的和
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

func rob(nums []int) int { // 同198.打家劫舍
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
