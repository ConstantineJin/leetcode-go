package main

import "slices"

func minCost(nums []int, x int) int64 {
	var n = len(nums)
	var s = make([]int, n) // s[i]统计操作i次的总成本
	for i := range s {
		s[i] = x * i
	}
	for i, num := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形）
			num = min(num, nums[j%n]) // 维护从nums[i]到nums[j]的最小值
			s[j-i] += num             // 累加操作j-i次的花费
		}
	}
	return int64(slices.Min(s))
}
