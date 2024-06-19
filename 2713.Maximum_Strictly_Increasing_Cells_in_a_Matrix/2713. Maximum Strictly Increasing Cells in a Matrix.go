package main

import "slices"

type pair struct{ x, y int }

// 定义 f[i][j] 表示到达 mat[i][j] 时，访问过的单元格的最大数量（包含 mat[i][j]）
// 那么答案就是所有 f[i][j] 的最大值
func maxIncreasingCells(mat [][]int) int {
	g := make(map[int][]pair) // key 是单元格值，value 是其坐标
	for i, row := range mat {
		for j, v := range row {
			g[v] = append(g[v], pair{i, j})
		}
	}
	keys := make([]int, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	slices.Sort(keys)                                                 // 按照单元格值升序排序
	rowMax, colMax := make([]int, len(mat)), make([]int, len(mat[0])) // 用一个长为 m 的数组 rowMax 维护每一行的 f 值的最大值，用一个长为 n 的数组 colMax 维护每一列的 f 值的最大值
	for _, key := range keys {                                        // 从小到大遍历所有的单元格值
		pos := g[key] // 当前单元格值的所有坐标
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先更新完 mx，再更新 rowMax 和 colMax
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return slices.Max(rowMax)
}
