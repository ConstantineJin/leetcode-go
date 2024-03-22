package main

import (
	"math"
	"sort"
)

func minimumVisitedCells(grid [][]int) (ans int) {
	var m, n = len(grid), len(grid[0])
	type pair struct{ x, i int }
	var colStacks = make([][]pair, n) // 每列的单调栈
	var rowSt []pair                  // 行单调栈
	for i := m - 1; i >= 0; i-- {
		rowSt = rowSt[:0] // 清空栈
		for j := n - 1; j >= 0; j-- {
			var colSt = colStacks[j]
			if i < m-1 || j < n-1 {
				ans = math.MaxInt
			}
			if g := grid[i][j]; g > 0 { // 可以向右/向下跳
				// 在单调栈上二分查找最优转移来源
				k := sort.Search(len(rowSt), func(k int) bool { return rowSt[k].i <= j+g })
				if k < len(rowSt) {
					ans = rowSt[k].x
				}
				k = sort.Search(len(colSt), func(k int) bool { return colSt[k].i <= i+g })
				if k < len(colSt) {
					ans = min(ans, colSt[k].x)
				}
			}
			if ans < math.MaxInt {
				ans++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(rowSt) > 0 && ans <= rowSt[len(rowSt)-1].x {
					rowSt = rowSt[:len(rowSt)-1]
				}
				rowSt = append(rowSt, pair{ans, j})
				for len(colSt) > 0 && ans <= colSt[len(colSt)-1].x {
					colSt = colSt[:len(colSt)-1]
				}
				colStacks[j] = append(colSt, pair{ans, i})
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return
}
