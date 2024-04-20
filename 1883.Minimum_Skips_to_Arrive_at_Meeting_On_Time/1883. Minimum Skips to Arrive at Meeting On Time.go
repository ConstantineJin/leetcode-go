package main

import (
	"math"
)

func minSkips(dist []int, speed int, hoursBefore int) int {
	var n = len(dist)
	var f = make([][]int, n+1) // f[i][j]表示经过了dist[0]到dist[i−1]的i段道路，并且跳过了j次的最短用时
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if i > j { // 如果i==j，不能不跳过
				f[i][j] = min(f[i][j], ((f[i-1][j]+dist[i-1]-1)/speed+1)*speed)
			}
			if j > 0 { // 如果j==0，不能跳过
				f[i][j] = min(f[i][j], f[i-1][j-1]+dist[i-1])
			}
		}
	}
	for j := 0; j <= n; j++ {
		if f[n][j] <= hoursBefore*speed {
			return j
		}
	}
	return -1
}
