package main

import "math"

func minFallingPathSum(grid [][]int) int {
	var n = len(grid)
	var prevFirstMin, prevSecondMin, prevFirstMinIdx = 0, 0, -1 // 只需三个变量记录每行最小的两个和（两者可能相等），以及最小和的index
	for i := 0; i < n; i++ {
		var curFirstMin, curSecondMin, curFirstMinIdx = math.MaxInt, math.MaxInt, -1
		for j := 0; j < n; j++ {
			var curSum = grid[i][j]
			if j == prevFirstMinIdx { // 如果当前的index恰等于上一行最小和的index，只能选择上一行第二小的和
				curSum += prevSecondMin
			} else {
				curSum += prevFirstMin
			}
			if curSum < curSecondMin { // 更新当前行最小和次小和，以及最小和的index
				curSecondMin = curSum
				if curSecondMin < curFirstMin {
					curFirstMin, curSecondMin = curSecondMin, curFirstMin
					curFirstMinIdx = j
				}
			}
		}
		prevFirstMin, prevSecondMin, prevFirstMinIdx = curFirstMin, curSecondMin, curFirstMinIdx // 覆盖
	}
	return prevFirstMin
}
