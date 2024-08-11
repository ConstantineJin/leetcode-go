package main

import "sort"

func maxUncrossedLines(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	mp := make(map[int][]int)
	for i, num := range nums2 {
		mp[num] = append(mp[num], i)
	}
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) (res int)
	dfs = func(i, j int) (res int) {
		if i == m {
			return
		}
		if mem[i][j+1] != -1 {
			return mem[i][j+1]
		}
		defer func() { mem[i][j+1] = res }()
		if indices, ok := mp[nums1[i]]; ok && len(indices) > 0 {
			k := sort.SearchInts(indices, j+1)
			if k == len(indices) {
				return dfs(i+1, j)
			} else {
				return max(dfs(i+1, j), dfs(i+1, indices[k])+1)
			}
		} else {
			return dfs(i+1, j)
		}
	}
	return dfs(0, -1)
}
