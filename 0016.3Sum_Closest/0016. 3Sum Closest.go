package main

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // nums升序排列
	var ans, n = nums[0] + nums[1] + nums[2], len(nums)
	for i, num := range nums[:n-2] {
		if i > 0 && num == nums[i-1] { // 如果当前最左数与前一个数相等，则后面的情况都已经计算过了
			continue
		}
		var sum = num + nums[i+1] + nums[i+2]
		if sum > target { // 固定最左数，三数之和最小值已经比target大，则后续不可能有更小的三数之和，判断是否需要更新ans后直接退出循环
			if sum-target < abs(ans-target) {
				ans = sum
			}
			break
		}
		sum = num + nums[n-1] + nums[n-2]
		if sum < target { // 固定最左数，三数之和最大值比target小，本次内循环可以直接跳过，继续下一次循环
			if target-sum < abs(ans-target) {
				ans = sum
			}
			continue
		}
		var j, k = i + 1, n - 1 // 中间数与最右数初始指向最左数的下一个数和数组末尾
		for j < k {             // 中右指针不相遇
			sum = num + nums[j] + nums[k]
			if abs(sum-target) < abs(ans-target) { // 如果当前三数之和更接近目标值
				ans = sum // 更新答案
			}
			if sum < target { // 三数之和比目标值小
				j++ // 向右移动中间的数
			} else if sum > target { // 三数之和比目标值大
				k-- // 向左移动最右的数
			} else { // 三数之和恰好等于sum
				return sum // 不可能有更接近的了，直接返回
			}
		}
	}
	return ans
}
