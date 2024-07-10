package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 两个元素做 OR 相当于对应的两个集合做并集
// 如果 a|b=a 说明 a 对应的集合是 b 对应的集合的超集
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, num := range nums { // 外循环正向遍历 nums
		ans = min(ans, abs(num-k)) // 单个元素也是一个子数组
		// 内循环从当前元素的前一个元素开始往前倒序遍历
		// 对于 k<j，nums[k] 一定是 nums[j] 的子集
		// 如果遇到做 OR 后不变，说明 nums[j] 开始往前的 nums[k] (0<=k<=j) 都是 nums[i] 的超集
		// 则后续循环不可能让值变小，直接退出内循环
		for j := i - 1; j >= 0 && nums[j]|num != nums[j]; j-- {
			nums[j] |= num // nums[j] 保存 nums[j:i+1] 的 OR
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}
