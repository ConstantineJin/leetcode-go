package main

import "strconv"

func findColumnWidth(grid [][]int) []int {
	var n = len(grid[0])
	var ans = make([]int, n)
	for _, row := range grid {
		for j, s := range row {
			if len(strconv.Itoa(s)) > ans[j] {
				ans[j] = len(strconv.Itoa(s))
			}
		}
	}
	return ans
}
