package main

import "sort"

func beautifulSubsets(nums []int, k int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	subset := make(map[int]int)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
		if subset[nums[i]-k] == 0 { // 选
			subset[nums[i]]++
			dfs(i + 1)
			subset[nums[i]]--
		}
		dfs(i + 1) // 不选
	}
	dfs(0)
	return ans - 1 // 排除空集
}
