package main

// 二分查找
//func searchMatrix(matrix [][]int, target int) bool {
//	for _, rows := range matrix {
//		i := sort.SearchInts(rows, target)
//		if i == len(matrix[0]) {
//			continue
//		}
//		if rows[i] == target {
//			return true
//		}
//	}
//	return false
//}

// Z形搜索
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1 // 从右上角开始
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target { // 任意一格的左值比当前格小，下值比当前格大。根据当前格子与target的大小关系决定移动方向
			y--
		} else {
			x++
		}
	}
	return false
}
