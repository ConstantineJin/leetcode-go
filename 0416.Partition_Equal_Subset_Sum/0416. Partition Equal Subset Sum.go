package main

import "sort"

func canPartition(nums []int) (ans bool) {
	sort.Ints(nums)
	var n = len(nums)
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 == 1 {
		return false
	}
	totalSum /= 2
	var mem = make([][]bool, n)
	for i := range mem {
		mem[i] = make([]bool, totalSum+1)
	}
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if i == n {
			return
		}
		if mem[i][sum] {
			return
		}
		defer func() { mem[i][sum] = true }()
		if sum+nums[i] < totalSum {
			dfs(i+1, sum)
			dfs(i+1, sum+nums[i])
		} else if sum+nums[i] > totalSum {
			return
		} else {
			ans = true
			return
		}
	}
	dfs(0, 0)
	return
}
