package main

import "slices"

func diagonalSort(mat [][]int) [][]int {
	var m, n = len(mat), len(mat[0])
	var arr = make([]int, min(m, n))
	for k := 1 - n; k < m; k++ { // k = i - j
		a := arr[:0]
		minI := max(k, 0)
		maxI := min(k+n, m)
		for i := minI; i < maxI; i++ {
			a = append(a, mat[i][i-k])
		}
		slices.Sort(a)
		for i := minI; i < maxI; i++ {
			mat[i][i-k] = a[i-minI]
		}
	}
	return mat
}
