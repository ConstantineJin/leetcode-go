package main

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var permute []int
	visit := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int{}, permute...))
			return
		}
		for j, num := range nums {
			if visit[j] || j > 0 && !visit[j-1] && num == nums[j-1] { // 对于重复出现的数字，只选第一个未访问过的
				continue
			}
			permute = append(permute, num)
			visit[j] = true
			dfs(i + 1)
			permute = permute[:len(permute)-1]
			visit[j] = false
		}
	}
	dfs(0)
	return
}
