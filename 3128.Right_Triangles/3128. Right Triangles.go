package main

func numberOfRightTriangles(grid [][]int) (ans int64) {
	rowCnt, colCnt := make([]int, len(grid)), make([]int, len(grid[0]))
	for i, row := range grid {
		for j, v := range row {
			rowCnt[i] += v
			colCnt[j] += v
		}
	}
	for i, row := range grid {
		for j, v := range row {
			ans += int64(v * max(rowCnt[i]-1, 0) * max(colCnt[j]-1)) // 枚举直角顶点，乘法原理
		}
	}
	return
}
