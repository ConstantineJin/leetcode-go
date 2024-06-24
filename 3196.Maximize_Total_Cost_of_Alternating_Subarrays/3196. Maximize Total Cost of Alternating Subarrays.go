package main

import "math"

func maximumTotalCost(nums []int) int64 {
	n := len(nums)
	mem := make([][2]int, n)
	for i := range mem {
		mem[i] = [2]int{math.MinInt, math.MinInt}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		p := &mem[i][j]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := dfs(i+1, 1) + nums[i] // 分割
		r := dfs(i+1, j^1)           // 不分割
		if j == 0 {
			r += nums[i]
		} else {
			r -= nums[i]
		}
		res = max(res, r)
		*p = res
		return res
	}
	return int64(dfs(0, 0))
}
