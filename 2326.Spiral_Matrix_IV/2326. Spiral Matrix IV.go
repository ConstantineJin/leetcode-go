package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrix(m, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}
	var x, y, i int // (x,y) 表示当前坐标，dirs[i] 表示当前前进方向
	for ; head != nil; head = head.Next {
		matrix[x][y] = head.Val
		newX, newY := x+dirs[i][0], y+dirs[i][1]
		if newX < 0 || newX == m || newY < 0 || newY == n || matrix[newX][newY] != -1 { // 如果遇到矩阵边界或遇到已经填写已填充的单元格则右转
			i = (i + 1) % 4
		}
		x, y = x+dirs[i][0], y+dirs[i][1]
	}
	return matrix
}
