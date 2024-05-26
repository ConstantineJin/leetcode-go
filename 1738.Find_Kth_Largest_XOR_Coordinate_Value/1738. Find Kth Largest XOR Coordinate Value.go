package main

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	xors := make([]int, 0, m*n)
	prefixXors := make([][]int, m+1)
	for i := range prefixXors {
		prefixXors[i] = make([]int, n+1)
	}
	for i, row := range matrix {
		for j, v := range row {
			prefixXors[i+1][j+1] = prefixXors[i+1][j] ^ prefixXors[i][j+1] ^ prefixXors[i][j] ^ v // 从(0,0)到(i,j)都被计算了三次（同一个数两次异或结果为0，任何数异或0结果不变），其余均异或了一次
		}
		xors = append(xors, prefixXors[i+1][1:]...)
	}
	sort.Ints(xors) // 也可使用快速选择算法
	return xors[len(xors)-k]
}
