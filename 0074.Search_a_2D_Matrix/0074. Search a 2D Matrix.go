package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	head := make([]int, m)
	for i, row := range matrix {
		head[i] = row[0]
	}
	i := sort.SearchInts(head, target) // 对每一行的头元素进行二分查找
	if i < m && matrix[i][0] == target {
		return true
	}
	if i == 0 {
		return false
	}
	i--
	j := sort.SearchInts(matrix[i], target) // 对元素可能存在的行进行二分查找
	if j < n && matrix[i][j] == target {
		return true
	}
	return false
}
