package main

import "sort"

// 动态规划，递归
//func maximumTotalDamage(power []int) int64 {
//	cnt := make(map[int]int)
//	for _, p := range power {
//		cnt[p]++
//	}
//	n := len(cnt)
//	arr := make([]int, 0, n)
//	for x := range cnt {
//		arr = append(arr, x)
//	}
//	sort.Ints(arr)
//	mem := make([]int, n)
//	for i := range mem {
//		mem[i] = -1
//	}
//	var dfs func(i int) int
//	dfs = func(i int) int {
//		if i == n {
//			return 0
//		}
//		if mem[i] > -1 {
//			return mem[i]
//		}
//		x := arr[i]
//		j := i
//		for j < n-1 && arr[j+1] <= x+2 {
//			j++
//		}
//		mem[i] = max(dfs(i+1), dfs(j+1)+x*cnt[x])
//		return mem[i]
//	}
//	return int64(dfs(0))
//}

// 动态规划，递推
func maximumTotalDamage(power []int) int64 {
	cnt := make(map[int]int) // key 是咒语，value 是该咒语的数量
	for _, p := range power {
		cnt[p]++
	}
	n := len(cnt)
	arr := make([]int, 0, n)
	for x := range cnt {
		arr = append(arr, x)
	}
	sort.Ints(arr)
	f := make([]int, n+1)
	var j int // 慢指针
	for i, x := range arr {
		for arr[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}
