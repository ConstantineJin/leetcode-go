package main

import "math"

const inf = math.MaxInt / 2

func countPaths(n int, roads [][]int) int {
	var g = make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, road := range roads {
		var x, y, d = road[0], road[1], road[2]
		g[x][y], g[y][x] = d, d
	}
	var dis = make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = inf
	}
	var f, done = make([]int, n), make([]bool, n)
	f[0] = 1
	for {
		var x = -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x == n-1 {
			return f[n-1]
		}
		done[x] = true
		for y, d := range g[x] {
			if dis[x]+d < dis[y] {
				dis[y], f[y] = dis[x]+d, f[x]
			} else if dis[x]+d == dis[y] {
				f[y] = (f[x] + f[y]) % (1e9 + 7)
			}
		}
	}
}
