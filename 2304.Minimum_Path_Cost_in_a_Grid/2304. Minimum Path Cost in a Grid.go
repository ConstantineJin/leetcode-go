package main

import (
	"math"
	"slices"
)

// 动态规划，递归
//func minPathCost(grid [][]int, moveCost [][]int) int {
//	m, n, ans := len(grid), len(grid[0]), math.MaxInt
//	mem := make([][]int, m) // 记忆化搜索
//	for i := range mem {
//		mem[i] = make([]int, n)
//	}
//	var dfs func(i, j int) int // dfs(i,j)表示从grid[i][j]出发到达最后一行的最小路径代价
//	dfs = func(i, j int) int {
//		if i == m-1 { // 递归边界
//			return grid[i][j]
//		}
//		p := &mem[i][j]
//		if *p != 0 { // 已经计算过了
//			return *p
//		}
//		res := math.MaxInt
//		for k, cost := range moveCost[grid[i][j]] {
//			res = min(res, dfs(i+1, k)+cost) // 移动到下一行第k列
//		}
//		*p = res + grid[i][j] // 保存结果
//		return *p
//	}
//	for j := 0; j < n; j++ {
//		ans = min(ans, dfs(0, j))
//	}
//	return ans
//}

// 递推
//func minPathCost(grid [][]int, moveCost [][]int) int {
//	m, n := len(grid), len(grid[0])
//	f := make([][]int, m)
//	for i := range f {
//		f[i] = make([]int, n)
//	}
//	f[m-1] = grid[m-1]
//	for i := m - 2; i >= 0; i-- { // i必须从下往上递推
//		for j, g := range grid[i] { // j的计算顺序无关紧要
//			f[i][j] = math.MaxInt
//			for k, c := range moveCost[g] { // 移动到下一行的第 k 列
//				f[i][j] = min(f[i][j], f[i+1][k]+c)
//			}
//			f[i][j] += g
//		}
//	}
//	return slices.Min(f[0])
//}

// 递推空间优化，结果直接保存在grid中
func minPathCost(grid [][]int, moveCost [][]int) int {
	for i := len(grid) - 2; i >= 0; i-- {
		for j, g := range grid[i] {
			res := math.MaxInt
			for k, c := range moveCost[g] {
				res = min(res, grid[i+1][k]+c)
			}
			grid[i][j] += res
		}
	}
	return slices.Min(grid[0])
}
