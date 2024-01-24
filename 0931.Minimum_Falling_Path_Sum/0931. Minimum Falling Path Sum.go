package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	var n, ans = len(matrix), math.MaxInt
	for i := n - 2; i >= 0; i-- { // 从倒数第二行开始向上递推
		matrix[i][0] += min(matrix[i+1][0], matrix[i+1][1]) // 每行首尾元素特殊处理
		matrix[i][n-1] += min(matrix[i+1][n-2], matrix[i+1][n-1])
		for j := 1; j < n-1; j++ {
			matrix[i][j] += min(matrix[i+1][j-1], matrix[i+1][j], matrix[i+1][j+1]) // 递推关系：每个元素的下降路径最小和等于当前元素的值加上左下、正下与右下元素的下降路径最小和的最小值
		}
	}
	for _, v := range matrix[0] { // 第一行的最小值就是答案
		ans = min(v, ans)
	}
	return ans
}
