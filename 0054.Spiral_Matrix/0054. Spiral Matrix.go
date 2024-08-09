package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	ans := make([]int, 0, m*n)
	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			ans = append(ans, matrix[top][j])
		}
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if left < right && top < bottom { // 当前层不是一维的
			for column := right - 1; column > left; column-- {
				ans = append(ans, matrix[bottom][column])
			}
			for row := bottom; row > top; row-- {
				ans = append(ans, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
