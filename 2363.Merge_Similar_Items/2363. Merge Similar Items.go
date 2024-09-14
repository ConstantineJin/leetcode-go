package main

import "slices"

func mergeSimilarItems(items1, items2 [][]int) [][]int {
	mp := make(map[int]int, len(items1)+len(items2))
	for _, item := range items1 {
		mp[item[0]] += item[1]
	}
	for _, item := range items2 {
		mp[item[0]] += item[1]
	}
	ans := make([][]int, 0, len(mp))
	for value, weight := range mp {
		ans = append(ans, []int{value, weight})
	}
	slices.SortFunc(ans, func(a, b []int) int { return a[0] - b[0] })
	return ans
}
