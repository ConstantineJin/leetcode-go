package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 两个元素做AND相当于对应的两个集合做交集
// 如果a&b=a说明a对应的集合是b对应的集合的子集
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, num := range nums { // 外循环正向遍历
		ans = min(ans, abs(num-k)) // 单个元素也是一个子数组
		// 内循环从当前元素的前一个元素开始往前倒序遍历
		// 对于k<j，nums[k]一定是nums[j]的子集
		// 如果遇到做AND后不变，说明nums[j]开始往前的nums[k](0<=k<=j)都是nums[i]的子集
		// 则后续循环不可能让值变小，直接退出内循环
		for j := i - 1; j >= 0 && nums[j]&num != nums[j]; j-- {
			nums[j] &= num // nums[j]保存nums[j:i+1]的AND
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}
