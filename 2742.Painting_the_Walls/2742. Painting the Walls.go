package main

import "math"

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
	}
	for i, c := range cost {
		t := time[i] + 1
		for j := n; j > 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}
