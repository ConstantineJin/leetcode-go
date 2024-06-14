package main

import "slices"

func findMaximumElegance(items [][]int, k int) int64 {
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] }) // 根据利润降序排序
	var totalProfit int
	visited := make(map[int]bool)
	var duplicate []int
	for _, item := range items[:k] {
		profit, category := item[0], item[1]
		totalProfit += profit
		if !visited[category] {
			visited[category] = true
		} else {
			duplicate = append(duplicate, profit)
		}
	}
	ans := totalProfit + len(visited)*len(visited)
	for _, item := range items[k:] {
		profit, category := item[0], item[1]
		if len(duplicate) > 0 && !visited[category] {
			visited[category] = true
			totalProfit += profit - duplicate[(len(duplicate)-1)]
			duplicate = duplicate[:(len(duplicate) - 1)]
		}
		ans = max(ans, totalProfit+len(visited)*len(visited))
	}
	return int64(ans)
}
