package main

type position struct{ x, y int }

func firstCompleteIndex(arr []int, mat [][]int) int {
	pos := make([]position, len(arr)+1) // 下标对应数字，值对应在mat中的坐标
	m, n := len(mat), len(mat[0])
	rows, cols := make([]int, m), make([]int, n) // 两个数组记录各行各列已填充元素的个数
	for i, row := range mat {
		for j, v := range row {
			pos[v] = position{x: i, y: j} // 遍历mat初始化pos
		}
	}
	for i, v := range arr { // 遍历arr根据pos
		x, y := pos[v].x, pos[v].y
		rows[x]++
		cols[y]++
		if rows[x] == n || cols[y] == m { // 如果行或者列被填满
			return i // 返回arr当前遍历到的下标
		}
	}
	return -1
}
