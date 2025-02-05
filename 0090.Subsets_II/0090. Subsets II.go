package main

import "slices"

func subsetsWithDup(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		num := nums[i]
		path = append(path, num)
		dfs(i + 1)
		path = path[:len(path)-1]
		for i++; i < n && nums[i] == num; i++ {
		}
		dfs(i)
	}
	dfs(0)
	return
}
